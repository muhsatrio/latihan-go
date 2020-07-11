package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// Create html from string

	name := "Muhammad Satrio Wicaksono"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` +
		name +
		`</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	// Create html from .gohtml file

	template, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err = os.Create("index2.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	err = template.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
