package main

import (
	"html/template"
	"log"
	"os"
)

type Restaurant struct {
	Name string
	Menu []Meal
}

type Item struct {
	Name  string
	Price float64
}

type Meal struct {
	MealName   string
	Appetizers []Item
	Entrees    []Item
	Drinks     []Item
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("blah.gohtml"))
}

func main() {

	b := Meal{
		MealName: "Breakfast",
		Entrees: []Item{
			Item{
				Name:  "Cereal",
				Price: 3.25,
			},
			Item{
				Name:  "Belgian Waffle",
				Price: 5.75,
			},
		},
		Drinks: []Item{
			Item{
				Name:  "Water",
				Price: 0,
			},
			Item{
				Name:  "Orange Juice",
				Price: 2.20,
			},
		},
	}
	l := Meal{
		MealName: "Lunch",
		Appetizers: []Item{
			Item{
				Name:  "Nachos",
				Price: 3.00,
			},
			Item{
				Name:  "Blooming Onion",
				Price: 5.20,
			},
		},
		Entrees: []Item{
			Item{
				Name:  "Suft n Turf",
				Price: 7.25,
			},
			Item{
				Name:  "Fish n Chips",
				Price: 7.75,
			},
		},
		Drinks: []Item{
			Item{
				Name:  "Water",
				Price: 0,
			},
			Item{
				Name:  "Soda",
				Price: 2.00,
			},
		},
	}
	d := Meal{
		MealName: "Dinner",
		Appetizers: []Item{
			Item{
				Name:  "Sliders",
				Price: 5.40,
			},
			Item{
				Name:  "Spiniach Artichoke Dip",
				Price: 6.80,
			},
		},
		Entrees: []Item{
			Item{
				Name:  "Tomahawk Steak",
				Price: 18.75,
			},
			Item{
				Name:  "Lobster Tail Ravioli",
				Price: 12.75,
			},
		},
		Drinks: []Item{
			Item{
				Name:  "Water",
				Price: 0,
			},
			Item{
				Name:  "Wine",
				Price: 6.00,
			},
		},
	}
	rest := Restaurant{
		Name: "Chez Josh",
		Menu: []Meal{b, l, d},
	}

	rests := struct {
		Restaurants []Restaurant
	}{
		Restaurants: []Restaurant{rest},
	}

	err := tpl.Execute(os.Stdout, rests)
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
