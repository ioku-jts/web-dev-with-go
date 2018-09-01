package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("someFile"))
}

func main() {
	tpl.Execute(os.Stdout, nil)
}
