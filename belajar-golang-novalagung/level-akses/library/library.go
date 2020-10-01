package library

import "fmt"

// fungsi ini memiliki akses public

// SayHello func
func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

// fungsi ini memiliki akses private
func introduce(name string) {
	fmt.Println("my name", name)
}

// Student struct
type Student struct {
	Name  string
	Grade int
}

// Murid var
var Murid = struct {
	Name  string
	Grade int
}{}

func init() {
	Murid.Name = "Muhammad Satrio Wicaksono"
	Murid.Grade = 24
}
