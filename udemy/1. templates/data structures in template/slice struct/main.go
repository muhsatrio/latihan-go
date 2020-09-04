package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type identitas struct {
	Nama   string
	Umur   int
	Zodiak string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	listPeople := []identitas{
		identitas{
			Nama:   "Satrio",
			Umur:   24,
			Zodiak: "Cancer",
		},
		identitas{
			Nama:   "Bambang",
			Umur:   22,
			Zodiak: "Taurus",
		},
		identitas{
			Nama:   "Setia",
			Umur:   25,
			Zodiak: "Aquarius",
		},
	}

	err := tpl.Execute(os.Stdout, listPeople)

	if err != nil {
		log.Fatalln(err)
	}
}
