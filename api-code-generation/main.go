package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	reMatchClass  = regexp.MustCompile(`^class (\S+) <([^>]+)>`)
	reMatchMethod = regexp.MustCompile(`^method (\S+) ([^(]+)\((.*)\)`)
	reMethodArg   = regexp.MustCompile(`(\S+)\s(\w+)`)
)

func main() {
	className := ""
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		switch {
		case reMatchClass.MatchString(line):
			className = generateClass(line)
		case reMatchMethod.MatchString(line):
			generateMethod(className, line)
		}
	}
}

func generateClass(raw string) string {
	matches := reMatchClass.FindStringSubmatch(raw)
	viewTpl.Lookup("package").Execute(os.Stdout, struct {
		PackageName string
		TypeName    string
	}{
		"api",
		matches[1],
	})

	return matches[1]
}

func generateMethod(className string, raw string) {
	matches := reMatchMethod.FindStringSubmatch(raw)

	if !isSimpleType(matches[1]) {
		// no support for non-simple type now
		return
	}

	returnType := convertJavaTypeToGoType(matches[1])
	methodName := matches[2]
	args, argsOk := parseArgs(matches[3])
	if !argsOk {
		// no support for non-simple type now
		return
	}

	// @TODO: fix me when #7 is ready
	returnType = ""

	err := viewTpl.Lookup("method").Execute(os.Stdout, struct {
		TypeName   string
		MethodName string
		ReturnType string
		Args       interface{}
	}{
		className,
		methodName,
		returnType,
		args,
	})

	if err != nil {
		panic(err)
	}
}

func parseArgs(raw string) (interface{}, bool) {
	if raw == "" {
		return nil, true
	}

	rawArgs := strings.Split(raw, "; ")

	args := make([]struct {
		Type string
		Name string
	}, len(rawArgs))

	for i, rawArg := range rawArgs {
		matches := reMethodArg.FindStringSubmatch(rawArg)
		if !isSimpleType(matches[1]) {
			return nil, false
		}

		args[i].Type = convertJavaTypeToGoType(matches[1])
		args[i].Name = matches[2]

	}

	return args, true
}

func convertJavaTypeToGoType(javaType string) string {
	switch javaType {
	case "void":
		return ""
	case "boolean":
		return "bool"
	case "Integer":
		fallthrough
	case "int":
		return "int"
	case "Double":
		fallthrough
	case "double":
		fallthrough
	case "float":
		return "float"
	case "String":
		fallthrough
	case "CharSequence":
		return "string"
	}

	return javaType
}

func isSimpleType(javaType string) bool {
	switch javaType {
	case "void":
		fallthrough
	case "boolean":
		fallthrough
	case "Integer":
		fallthrough
	case "int":
		fallthrough
	case "Double":
		fallthrough
	case "double":
		fallthrough
	case "float":
		fallthrough
	case "String":
		fallthrough
	case "CharSequence":
		return true
	}

	return false
}
