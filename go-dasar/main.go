package main

import (
	"fmt"
	"math"
)

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
	const pi float64 = 22.7
	var jariJari float64 = 10
	fmt.Printf("Luas lingkaran: %f\n", pi*math.Pow(jariJari, 2))

	// Operator Perbandingan
	fmt.Printf("2 > 5 : %t\n", 2 > 5)
	fmt.Printf("true and false: %t\n", true && false)

	// Penggunaan seleksi kondisi
	var point = 100.0
	if point > 100 {
		fmt.Println("Lebih besar dari 100")
	} else {
		fmt.Println("Lebih kecil dari 100")
	}

	// Penggunaan seleksi kondisi temporary
	if percent := point / 100; percent > 1 {
		fmt.Println("Lebih besar dari 1%")
	} else if percent >= 0.5 {
		fmt.Println("Lebih besar dari 0,5%")
	} else {
		fmt.Println("Lebih kecil dari 0.5%")
	}

	// Penggunaan switch case
	var nilai = 9
	switch nilai {
	case 10:
		fmt.Println("A")
	case 8, 9:
		fmt.Println("AB")
		fallthrough
		// fallthrough digunakan untuk memaksa kondisi dalam case tersebut dieksekusi, meskipun sebeenarnya case sebelumnya sudah ada yang bernilai true
	case 6, 7:
		fmt.Println("B")
	default:
		{
			fmt.Println("C")
		}
	}

	// Penggunaan switch case
	nilai = 9
	switch {
	case nilai <= 10 && nilai >= 8:
		fmt.Println("A")
	case nilai >= 8:
		fmt.Println("AB")
		fallthrough
		// fallthrough digunakan untuk memaksa kondisi dalam case tersebut dieksekusi, meskipun sebeenarnya case sebelumnya sudah ada yang bernilai true
	case nilai >= 6 && nilai <= 7:
		fmt.Println("B")
	default:
		{
			fmt.Println("C")
		}
	}
}
