package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	//parseGlob allows pattern filtering
	//tpl.Execute(os.Stdout, nil)
	tpl.ExecuteTemplate(os.Stdout, "template1.gohtml", "blah")
	tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", "bloop")
	tpl.ExecuteTemplate(os.Stdout, "template3.gohtml", nil)

	//if we wanted to create a new file
	newFile, err := os.Create("someFile")
	if err != nil {
		log.Fatalln("Error creating file.", err)
	}
	defer newFile.Close()

	err2 := tpl.ExecuteTemplate(newFile, "template3.gohtml", "new file!")
	if err2 != nil {
		log.Fatalln("Error executing template.", err)
	}
}
