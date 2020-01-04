package main

import "fmt"

func main() {
	name := "Muhammad Satrio Wicaksono"

	tpl := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `
	</body>
	</html>
	`
	fmt.Println(tpl)
}
