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

	// Perulangan dengan for standar
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Penggunaan for yang seperti while
	var i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Penggunaan for yang selalu bernilai true
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	for j := 1; j <= 5; j++ {
		for k := 0; k < j; k++ {
			fmt.Print('0')
		}
		fmt.Println()
	}

	// Inisialisasi nilai awal array
	var nama = [3]string{"Muhammad", "Satrio", "Wicaksono"}
	fmt.Println("Jumlah array:", len(nama))
	fmt.Println("Isi array:", nama)

	// Inisialisasi nilai awal array secara vertikal
	var buah = [3]string{
		"pepaya",
		"pisang",
		"jambu",
	}
	fmt.Println("Buah buahan:", buah)

	// Inisialisasi awal array dengan jumlah otomatis

	var sayur = [...]string{
		"bayam",
		"kol",
		"tomat",
	}
	fmt.Println("Sayur:", sayur)

	// Deklarasi variabel multidimensi
	var grid = [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(grid)

	// Iterasi array menggunakan len
	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}

	// Iterias array menggunakan range
	for i, syr := range sayur {
		fmt.Println(i, syr)
	}

	// Iterasi array menggunakan range tanpa number
	for _, syr := range sayur {
		fmt.Println(syr)
	}

	// Penggunaan make untuk alokasi array
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"
	fmt.Println(fruits)
}
