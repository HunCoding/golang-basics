package main

import (
	"os"
	"text/template"
)

func main() {
	templ := `package {{.Package}}

import "fmt"

type {{.StructName}} struct {
{{range .Fields}}	{{.Name}} {{.Type}}
{{end}}
}

func printStruct(s {{.StructName}}) {
   fmt.Println(s)
}

func main() {
	s := {{.StructName}}{
		{{range .Fields}}		{{.Name}}: {{.Value}},
		{{end}}
	}

	printStruct(s)
}
	`

	type Field struct {
		Name  string
		Type  string
		Value string
	}

	type Model struct {
		Package    string
		StructName string
		Fields     []Field
	}

	data := Model{
		Package:    "main",
		StructName: "Person",
		Fields: []Field{
			{Name: "Name", Type: "string", Value: `"John"`},
			{Name: "Youtube", Type: "string", Value: `"youtube.com/@huncoding"`},
		},
	}

	t := template.Must(template.New("main").Parse(templ))
	t.Execute(os.Stdout, data)
}
