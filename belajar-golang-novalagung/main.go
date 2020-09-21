package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = min + rand.Intn(max-min)
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider, %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}

func calculateAverage(numbers ...int) float64 {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func myHobbies(name string, hobbies ...string) {
	var convertHobbiesToString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is %s\n", name)
	fmt.Printf("My hobbies are: %s\n", convertHobbiesToString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

// FilterCallback type
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func change(original *int, value int) {
	*original = value
}

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

	fmt.Println(sayur[0:2])

	// Deklarasi variabel multidimensi
	var grid = [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(grid)

	// Iterasi array menggunakan len
	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}

	// Iterasi array menggunakan range
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

	// Penggunaan slice, slice tidak perlu mendefinisikan jumlah elemen pada awal deklarasi

	var vehicles = []string{"mobil", "motor", "perahu", "pesawat"}

	fmt.Println(vehicles[0:2]) //dimulai dari indeks-0, sampai SEBELUM indeks-2
	fmt.Println(vehicles[2:])  //semua elemen mulai dari indeks ke-2
	fmt.Println(vehicles[:2])  //semua elemen sampai indeks sebelum ke-2

	// fungsi len untuk menghitung panjang array
	fmt.Println("Length:", len(vehicles))

	// fungsi len untuk menghitung kapasitas array
	vehiclesNew := vehicles[0:2]
	fmt.Println("Length:", len(vehiclesNew))
	fmt.Println("Capacity:", cap(vehiclesNew))
	fmt.Println("Capacity:", cap(vehicles[1:3]))

	// fungsi append untuk menambah element pada slice
	vehicles = append(vehicles, "sepeda")
	fmt.Println(vehicles)

	// Penggunaan map
	var chicken = map[string]int{}
	chicken["test"] = 50
	fmt.Println(chicken["test"])

	// Definisi map di awal
	var identitas = map[string]string{
		"firstName": "Muhammad",
		"lastName":  "Satrio",
		"NIM":       "1301150052",
		"Kelas":     "IF-39-10",
	}
	fmt.Println(identitas)

	// Iterasi menggunakan map
	for key, val := range identitas {
		fmt.Println(key, ": ", val)
	}

	// Menghapus nama terakhir di map identitas
	delete(identitas, "lastName")
	for key, val := range identitas {
		fmt.Println(key, ": ", val)
	}

	// Cek apakah data kelas ada di map
	var _, isExist = identitas["Kelas"]

	if isExist {
		fmt.Println("Ada")
	} else {
		fmt.Println("Tidak ada")
	}

	// Slice dikombinasikan dengan map

	var mahasiswa = []map[string]string{
		{
			"firstName": "Muhammad",
			"lastName":  "Satrio",
		},
		{
			"firstName": "Haya",
			"lastName":  "Majidatul",
		},
	}
	fmt.Println(mahasiswa)

	// Penggunaan fungsi

	var names = []string{
		"Muhammad",
		"Satrio",
	}
	printMessage("sampurasun", names)

	// Penggunaan fungsi dengan return

	rand.Seed(time.Now().Unix())

	var randomValue int

	randomValue = randomWithRange(2, 20)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	// Penggunaan return untuk menghentikan proses dalam fungsi

	divideNumber(10, 2)
	divideNumber(4, 0)

	// Penggunaan fungsi multiple return

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran: %.2f\n", area)
	fmt.Printf("keliling lingkaran: %.2f\n", circumference)

	// Penggunaan fungsi variadic: fungsi dengan parameter sejenis yang tak terbatas

	var avg = calculateAverage(1, 2, 3, 4, 5)
	var numbers = []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	fmt.Printf("Rata-rata: %.2f\n", avg)
	fmt.Printf("Rata-rata: %.2f\n", calculateAverage(numbers...)) //fungsi variadic dengan data slice
	var listHobbies = []string{
		"makan",
		"minum",
		"tidur",
	}
	myHobbies("satrio", listHobbies...)

	// Penggunaan fungsi closure: fungsi yang disimpan di dalam variabel

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %d, max: %d\n", numbers, min, max)

	// Penggunaan fungsi closure yang bisa dieksekusi langsung saat deklarasinya

	numbers = []int{1, 2, 3, 4, 5}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				r = append(r, e)
			}
		}
		return r
	}(3)

	fmt.Printf("Original numbers: %v\n", numbers)
	fmt.Printf("New numbers: %v\n", newNumbers)

	// Penggunaam fungsi closure sebagai return

	var lengthNum, getNumbers = findMax(numbers, 3)

	fmt.Println(lengthNum)
	fmt.Println(getNumbers())

	// Penerapan fungsi sebagai parameter

	data := []string{"muhammad", "satrio", "wicaksono"}
	var dataContainsW = filter(data, func(each string) bool {
		return strings.Contains(each, "w")
	})
	var dataLength6 = filter(data, func(each string) bool {
		return len(each) == 6
	})
	fmt.Println(dataContainsW)
	fmt.Println(dataLength6)

	// Penerapan pointer

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA value: ", numberA)
	fmt.Println("numberA address: ", &numberA)

	fmt.Println("numberB value: ", *numberB)
	fmt.Println("numberB address: ", numberB)

	numberA = 8

	fmt.Println("numberA value: ", numberA)
	fmt.Println("numberA address: ", &numberA)

	fmt.Println("numberB value: ", *numberB)
	fmt.Println("numberB address: ", numberB)

	*numberB = 10

	fmt.Println("numberA value: ", numberA)
	fmt.Println("numberA address: ", &numberA)

	fmt.Println("numberB value: ", *numberB)
	fmt.Println("numberB address: ", numberB)

	// Pointer sebagai parameter dalam fungsi

	var number = 4

	fmt.Println("Before: ", number)

	change(&number, 10)

	fmt.Println("After: ", number)

}
