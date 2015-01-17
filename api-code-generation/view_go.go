package main

import (
	"strings"
	"text/template"

	"github.com/seletskiy/tplutil"
)

var viewTpl = template.New("go")

var _ = template.Must(viewTpl.New("package").Parse(strings.TrimLeft(`
package {{.PackageName}}

type {{.TypeName}} struct {
	id string
	server *Server
}

`, "\n")))

var _ = template.Must(viewTpl.New("args").Funcs(tplutil.Last).Parse(
	tplutil.Strip(`
	{{range $i, $_ := .Args}}
		{{.Name}} {{.Type}}
		{{if last $i $.Args | not}}, {{end}}
	{{end}}
`)))

var _ = template.Must(viewTpl.New("args_in_call").Funcs(tplutil.Last).Parse(
	tplutil.Strip(`
	{{range $i, $_ := .Args}}
		{{if eq .Type "float"}}
			float({{.Name}})
		{{else}}
			{{.Name}}
		{{end}}
		{{", "}}
	{{end}}
`)))

var _ = template.Must(viewTpl.New("return").Funcs(tplutil.Last).Parse(
	tplutil.Strip(`
	{{" "}}{{if .ReturnType}}{{.ReturnType}} {{end}}
`)))

var _ = template.Must(viewTpl.New("method").Parse(strings.TrimLeft(`
func (obj {{.TypeName}}) {{.MethodName}}({{template "args" .}}){{template "return" .}}{
	obj.server.Call(map[string]interface{}{
		obj.id,
		"{{.TypeName}}",
		"{{.MethodName}}",
		{{template "args_in_call" .}}
	})
}

`, "\n")))
