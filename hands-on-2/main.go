package main

import (
	"html/template"
	"log"
	"os"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("someFile"))
}

func main() {
	h := Hotel{
		Name:    "Lair",
		Address: "123 Blah St",
		City:    "Blooptown",
		Zip:     12345,
		Region:  "Southern",
	}
	e := Hotel{
		Name:    "Phils House",
		Address: "123 Blah St",
		City:    "Blooptown",
		Zip:     12345,
		Region:  "Central",
	}
	l := Hotel{
		Name:    "Kevins Kot",
		Address: "123 Blah St",
		City:    "Blooptown",
		Zip:     12345,
		Region:  "Southern",
	}
	p := Hotel{
		Name:    "Brendans Bank",
		Address: "123 Blah St",
		City:    "Blooptown",
		Zip:     12345,
		Region:  "Northern",
	}

	hotels := []Hotel{h, e, l, p}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln("Error", err)
	}
}
