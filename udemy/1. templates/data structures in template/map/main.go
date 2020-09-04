package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	tempatTinggal := map[string]string{
		"Satrio":  "Bekasi",
		"Bambang": "Makassar",
		"Slamet":  "Jogja",
	}

	err := tpl.Execute(os.Stdout, tempatTinggal)

	if err != nil {
		log.Fatalln(err)
	}

}
