package main

import (
	"fmt"
	"latihan-go/belajar-golang-novalagung/level-akses/library"
)

func main() {
	library.SayHello("satrio")
	var s1 = library.Student{Name: "satrio", Grade: 24}
	fmt.Println(s1.Name)
	fmt.Println(s1.Grade)

	fmt.Println(library.Murid.Name)
	fmt.Println(library.Murid.Grade)
}
