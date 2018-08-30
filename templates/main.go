package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatalln("Error parsing files.", err)
	}
	//tpl.Execute(os.Stdout, nil)
	tpl.ExecuteTemplate(os.Stdout, "template1.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "template3.gohtml", nil)
}
