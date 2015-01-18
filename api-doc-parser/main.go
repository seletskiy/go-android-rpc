package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/docopt/docopt-go"
	"golang.org/x/net/html"

	"log"

	"launchpad.net/xmlpath"
)

var usage = `Extract Android API methods list for later code generation.

Tool will extract all classes with theirs public methods signatures from
official android documentation and will output it in following format:

    class {class_name} <{url}>
    method {return_type} {method_name}({method_arg}; {method_arg}; ...)
    method {return_type} {method_name}({method_arg}; {method_arg}; ...)
    ...

Meant to be pipeable into code generation tool.

To output only specific class from package use:
  $0 -g '^CLASSNAME$' PACKAGE

Usage:
  $0 [options] <package_name>
  $0 -h|--help

Options:
  -g=<re>    Output only class names matching regular expression
             [default: ^].
  -e         Output "non-base" classes too. E.g. include classes like
             "AbsListView.OnScrollListener".
  -v         Be verbose.
  -h --help  Show this help.
`

type ApiClass struct {
	Url     string
	Name    string
	Methods ApiMethods
}

type ApiMethod struct {
	Name       string
	ReturnType string
	MethodArgs ApiMethodArgs
}

type ApiMethodArgs []ApiMethodArg
type ApiMethods []ApiMethod

type ApiMethodArg struct {
	Type string
	Name string
}

var (
	xpathApiClassA       = xmlpath.MustCompile(`//td[@class='jd-linkcol']/a`)
	xpathApiClassAHref   = xmlpath.MustCompile(`@href`)
	xpathPubApiMethodsTr = xmlpath.MustCompile(
		`//table[@id='pubmethods']//tr`,
	)
)

var (
	xpathPubApiMethodsReturnType         = xmlpath.MustCompile(`td[1]`)
	xpathPubApiMethodsApiMethodSignature = xmlpath.MustCompile(`td[2]/nobr`)
)

var (
	reApiMethodApiMethodArgs = regexp.MustCompile(`(\w+)\((.*)\)`)

	// some api pages contains unescaped html sequences, like like list<map<map>>
	reBrokenApiClassDescription = regexp.MustCompile(`\w+<([\w.,\s<>]+)>`)

	// \xa0 stands for non-breakable space, we're parsing html after all
	reSpaces               = regexp.MustCompile(`\s+|\xa0`)
	reMatchApiApiMethodArg = regexp.MustCompile(`(\w+|\w+<.*>)\s+(\w+)`)
)

var (
	baseUrl      = `https://developer.android.com/%s`
	packagesPath = `reference/android/%s/package-summary.html`
)

var debug = false

func main() {
	args, _ := docopt.Parse(
		strings.Replace(usage, `$0`, os.Args[0], -1),
		nil, true, "v1", false,
	)

	packageName := strings.Replace(
		args["<package_name>"].(string), `android.`, ``, -1,
	)

	debug = args["-v"].(bool)

	classes := getClassesList(packageName)

	nameRegexp := regexp.MustCompile(args["-g"].(string))

	baseOnly := !args["-e"].(bool)
	for _, class := range classes {
		if baseOnly && strings.Contains(class.Name, ".") {
			continue
		}

		if !nameRegexp.MatchString(class.Name) {
			continue
		}

		extractMethodsList(&class)

		fmt.Println(class)
	}
}

func getClassesList(packageName string) []ApiClass {
	root := getXpathParsedHTML(
		fmt.Sprintf(baseUrl, fmt.Sprintf(packagesPath, packageName)),
	)

	classes := []ApiClass{}

	linksIterator := xpathApiClassA.Iter(root)
	for linksIterator.Next() {
		node := linksIterator.Node()
		href, _ := xpathApiClassAHref.String(node)
		classes = append(classes, ApiClass{
			Url:  fmt.Sprintf(baseUrl, href),
			Name: node.String(),
		})
	}

	return classes
}

func extractMethodsList(class *ApiClass) {
	root := getXpathParsedHTML(class.Url)

	methodsIterator := xpathPubApiMethodsTr.Iter(root)

	for methodsIterator.Next() {
		node := methodsIterator.Node()

		returnTypeRaw, ok := xpathPubApiMethodsReturnType.String(node)
		if !ok {
			continue
		}

		methodSignatureRaw, _ := xpathPubApiMethodsApiMethodSignature.String(
			node,
		)

		name, args := parseSignature(methodSignatureRaw)
		method := ApiMethod{
			ReturnType: parseReturnType(returnTypeRaw),
			Name:       name,
			MethodArgs: args,
		}

		class.Methods = append(class.Methods, method)
	}
}

func fixHtml(input io.Reader) io.Reader {
	htmlBytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	htmlString := string(htmlBytes)

	brokenMatches := reBrokenApiClassDescription.FindAllStringSubmatch(
		htmlString, -1,
	)

	for _, brokenMatch := range brokenMatches {
		fixedMatch := brokenMatch[0]
		fixedMatch = strings.Replace(fixedMatch, `<`, `&lt;`, -1)
		fixedMatch = strings.Replace(fixedMatch, `>`, `&gt;`, -1)

		htmlString = strings.Replace(
			htmlString, brokenMatch[0], fixedMatch, -1,
		)
	}

	root, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}

	htmlBuffer := &bytes.Buffer{}
	html.Render(htmlBuffer, root)

	return htmlBuffer
}

func parseReturnType(raw string) string {
	returnType := strings.TrimSpace(reSpaces.ReplaceAllString(raw, ` `))

	cleanedUpType := returnType
	cleanedUpType = strings.Replace(returnType, `abstract `, ``, 1)
	cleanedUpType = strings.Replace(returnType, `final `, ``, 1)

	return cleanedUpType
}

func parseSignature(raw string) (name string, args []ApiMethodArg) {
	raw = reSpaces.ReplaceAllLiteralString(raw, ` `)

	splittedRaw := reApiMethodApiMethodArgs.FindStringSubmatch(raw)

	name = strings.TrimSpace(splittedRaw[1])

	argsRaw := reMatchApiApiMethodArg.FindAllStringSubmatch(splittedRaw[2], -1)

	args = []ApiMethodArg{}
	for _, argRaw := range argsRaw {
		args = append(args, ApiMethodArg{
			Type: argRaw[1],
			// "_" for possible intersections with go keywords
			Name: argRaw[2] + "_",
		})
	}

	return name, args
}

func getXpathParsedHTML(url string) *xmlpath.Node {
	if debug {
		log.Printf("GET %s", url)
	}

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fixedHtml := fixHtml(response.Body)
	root, err := xmlpath.ParseHTML(fixedHtml)
	if err != nil {
		panic(err)
	}

	return root
}

func (class ApiClass) String() string {
	return fmt.Sprintf(
		"class %s <%s>\n%s", class.Name, class.Url, class.Methods,
	)
}

func (args ApiMethods) String() string {
	result := make([]string, len(args))
	for i, arg := range args {
		result[i] = fmt.Sprint(arg)
	}

	return strings.Join(result, "\n")
}

func (method ApiMethod) String() string {
	return fmt.Sprintf(
		"method %s %s(%s)", method.ReturnType, method.Name, method.MethodArgs,
	)
}

func (args ApiMethodArgs) String() string {
	result := make([]string, len(args))
	for i, arg := range args {
		result[i] = fmt.Sprint(arg)
	}

	// ; here for easy split (, can be seen in complex types like map<a, b>)
	return strings.Join(result, "; ")
}

func (arg ApiMethodArg) String() string {
	return fmt.Sprintf("%s %s", arg.Type, arg.Name)
}
