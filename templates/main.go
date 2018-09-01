package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	//parseGlob allows pattern filtering
	//tpl.Execute(os.Stdout, nil)
	ss := []string{"Josh", "Jay", "Jake", "Jack"}
	tpl.ExecuteTemplate(os.Stdout, "template1.gohtml", "blah")
	tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", ss)
	//tpl.ExecuteTemplate(os.Stdout, "template3.gohtml", nil)

	//if we wanted to create a new file
	// newFile, err := os.Create("someFile")
	// if err != nil {
	// 	log.Fatalln("Error creating file.", err)
	// }
	// defer newFile.Close()

	// err2 := tpl.ExecuteTemplate(newFile, "template3.gohtml", "new file!")
	// if err2 != nil {
	// 	log.Fatalln("Error executing template.", err)
	// }
}
