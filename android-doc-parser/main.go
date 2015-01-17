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

	"golang.org/x/net/html"

	"log"

	"launchpad.net/xmlpath"
)

type APIClass struct {
	URL        string
	Name       string
	APIMethods []APIMethod
}

type APIMethod struct {
	Name          string
	ReturnType    string
	APIMethodArgs APIMethodArgs
}

type APIMethodArgs []APIMethodArg

type APIMethodArg struct {
	Type string
	Name string
}

var (
	xpathAPIClassA       = xmlpath.MustCompile(`//td[@class='jd-linkcol']/a`)
	xpathAPIClassAHref   = xmlpath.MustCompile(`@href`)
	xpathPubAPIMethodsTr = xmlpath.MustCompile(
		`//table[@id='pubmethods']//tr`,
	)
)

var (
	xpathPubAPIMethodsReturnType         = xmlpath.MustCompile(`td[1]`)
	xpathPubAPIMethodsAPIMethodSignature = xmlpath.MustCompile(`td[2]/nobr`)
)

var (
	reAPIMethodAPIMethodArgs = regexp.MustCompile(`(\w+)\((.*)\)`)

	// some api pages contains unescaped html sequences, like like list<map<map>>
	reBrokenAPIClassDescription = regexp.MustCompile(`\w+<([\w.,\s<>]+)>`)

	// \xa0 stands for non-breakable space, we're parsing html after all
	reSpaces               = regexp.MustCompile(`\s+|\xa0`)
	reMatchAPIAPIMethodArg = regexp.MustCompile(`(\w+|\w+<.*>)\s+(\w+)`)
)

var (
	baseUrl      = `https://developer.android.com/%s`
	packagesPath = `/reference/android/%s/package-summary.html`
)

func main() {
	reader, err := os.Open("package-summary.html")
	if err != nil {
		panic(err)
	}

	root, err := xmlpath.ParseHTML(reader)
	if err != nil {
		panic(err)
	}

	classes := []APIClass{}

	linksIterator := xpathAPIClassA.Iter(root)
	for linksIterator.Next() {
		node := linksIterator.Node()
		href, _ := xpathAPIClassAHref.String(node)
		classes = append(classes, APIClass{
			URL:  fmt.Sprintf(baseUrl, href),
			Name: node.String(),
		})
	}

	//classes = []APIClass{
	//    APIClass{
	//        URL: "https://developer.android.com//reference/android/widget/SimpleExpandableListAdapter.html",
	//    },
	//}

	for _, class := range classes {
		log.Printf("%s", class)

		response, err := http.Get(class.URL)
		if err != nil {
			panic(err)
		}

		fixedHtml := fixHtml(response.Body)
		root, err := xmlpath.ParseHTML(fixedHtml)
		if err != nil {
			fmt.Println(fixedHtml)
			panic(err)
		}

		methodsIterator := xpathPubAPIMethodsTr.Iter(root)

		for methodsIterator.Next() {
			node := methodsIterator.Node()

			returnTypeRaw, ok := xpathPubAPIMethodsReturnType.String(node)
			if !ok {
				continue
			}

			methodSignatureRaw, _ := xpathPubAPIMethodsAPIMethodSignature.String(node)

			name, args := parseSignature(methodSignatureRaw)
			method := APIMethod{
				ReturnType:    parseReturnType(returnTypeRaw),
				Name:          name,
				APIMethodArgs: args,
			}

			log.Printf("%s", method)
		}
	}
}

func fixHtml(input io.Reader) io.Reader {
	htmlBytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	htmlString := string(htmlBytes)

	brokenMatches := reBrokenAPIClassDescription.FindAllStringSubmatch(
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

	return strings.Replace(returnType, `abstract `, ``, 1)
}

func parseSignature(raw string) (name string, args []APIMethodArg) {
	raw = reSpaces.ReplaceAllLiteralString(raw, ` `)

	splittedRaw := reAPIMethodAPIMethodArgs.FindStringSubmatch(raw)

	name = strings.TrimSpace(splittedRaw[1])

	argsRaw := reMatchAPIAPIMethodArg.FindAllStringSubmatch(splittedRaw[2], -1)

	args = []APIMethodArg{}
	for _, argRaw := range argsRaw {
		args = append(args, APIMethodArg{
			Type: argRaw[1],
			Name: argRaw[2],
		})
	}

	return name, args
}

func (class APIClass) String() string {
	return fmt.Sprintf("class %s <%s>", class.Name, class.URL)
}

func (method APIMethod) String() string {
	return fmt.Sprintf(
		"method %s %s(%s)", method.ReturnType, method.Name, method.APIMethodArgs,
	)
}

func (args APIMethodArgs) String() string {
	result := make([]string, len(args))
	for i, arg := range args {
		result[i] = fmt.Sprint(arg)
	}

	return strings.Join(result, ", ")
}

func (arg APIMethodArg) String() string {
	return fmt.Sprintf("%s %s", arg.Type, arg.Name)
}
