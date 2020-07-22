package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type identity struct {
	Nama   string
	Umur   int
	Zodiak string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	person := identity{
		Nama:   "Satrio",
		Umur:   24,
		Zodiak: "Cancer",
	}

	err := tpl.Execute(os.Stdout, person)

	if err != nil {
		log.Fatalln(err)
	}

}
