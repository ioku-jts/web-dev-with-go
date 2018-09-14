package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func init() {

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "blank")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Josh Shih")
}

func cat(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln("Error parsing file.", err)
	}

	cats := []string{"Chowder", "Momo"}

	err = tpl.ExecuteTemplate(w, "template.gohtml", cats)

	if err != nil {
		log.Fatalln("Error executing template.", err)
	}
}
