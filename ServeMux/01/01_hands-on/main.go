package main

import (
	"io"
	"net/http"
)

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
	io.WriteString(w, "Chowder")
}
