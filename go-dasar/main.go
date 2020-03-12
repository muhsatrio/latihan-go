package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// dibawah ini output dengan beberapa value
	fmt.Println("nama", "saya", "satrio")

	// Deklarasi variabel dengan var
	var firstName string = "Muhammad"
	var middleName string
	middleName = "Satrio"

	fmt.Println(firstName, middleName)

	// Output dengan Printf
	fmt.Printf("%s %s\n", firstName, middleName)

	// Deklarasi variabel dengan inference
	// inference hanya bisa dilakukan didalam fungsi, tidak bisa dilakukan di luar fungsi / global
	lastName := "Wicaksono"
	fmt.Println(lastName)

	// Deklarasi multi variabel
	firstName, middleName, lastName = "Haya", "Majidatul", "Khasna"
	fmt.Println(firstName, middleName, lastName)

	// Deklarasi variabel dump
	satu, _ := "saya", "ganteng"
	fmt.Println(satu)

	// Deklarasi variabel pointer
	dua := new(int)
	fmt.Println(dua)
	fmt.Println(*dua)

	// Penggunaan const
	const pi float32 = 22.7
	var jariJari float32 = 10
	fmt.Printf("Luas lingkaran: %f\n", pi*jariJari*jariJari)
}
