package main

import (
	"fmt"
)

func main() {
	name := "Muhammad Satrio Wicaksono"

	tpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
		</head>
		<body>
			h1` + name + `
		</body>
		</html>
	`
	fmt.Println(tpl)
}
