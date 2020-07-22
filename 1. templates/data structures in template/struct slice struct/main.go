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

type kelas struct {
	Nama      string
	Kapasitas int
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

	listKelas := []kelas{
		kelas{
			Nama:      "IF-39-10",
			Kapasitas: 30,
		},
		kelas{
			Nama:      "IF-39-11",
			Kapasitas: 31,
		}, kelas{
			"IF-39-12",
			32,
		},
	}

	data := struct {
		Identitas []identitas
		Kelas     []kelas
	}{
		listPeople,
		listKelas,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
