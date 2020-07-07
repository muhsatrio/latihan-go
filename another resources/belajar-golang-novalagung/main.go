package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Ini perintah untuk menuliskan hello world
	fmt.Println("hello world")

	// ini deklarasi manifest typing dengan nilai variabel yang langsung diisi
	var firstName string = "Muhammad"

	// ini deklarasi manifest typing dengan nilai variabel tidak langsung diisi
	var middleName string
	middleName = "Satrio"

	fmt.Printf("Halo %s %s \n", firstName, middleName)

	// deklarasi type interference
	lastName := "Wicaksono"
	var message = "Nice to meet you"

	fmt.Printf("Halo %s %s %s %s\n", firstName, middleName, lastName, message)
	// Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.

	// Deklarasi multi variabel
	var alamatBlog, alamatWebsite = "blog.msatrio.com", "msatrio.com"
	fmt.Printf("Alamat blog: %s %s\n", alamatBlog, alamatWebsite)

	// Deklarasi multi variabel beda tipe data
	var angka, huruf = 12, "satrio"
	angka2, huruf2 := 24, "wicaksono"
	fmt.Printf("%d %d %s %s\n", angka, angka2, huruf, huruf2)

	// Variabel underscore
	angka, _ = 35, 45
	fmt.Println(angka)

	// Type of variable
	var angka3 = 123456789101112
	fmt.Println(reflect.TypeOf(angka3).Kind())

	var inputan string
	fmt.Scan(inputan)
	fmt.Println(inputan)

}
