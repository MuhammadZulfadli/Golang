package main

import "fmt"

func main() {
	// inisialisasi tipe data terlebih dahulu
	var nama string

	nama = "Muhammad Zulfadli"

	fmt.Println(nama)

	nama = "Muhammad Zulfadli Simatupang"

	fmt.Println(nama)

	// Ketika tidak di inisialisasi
	var myGirlName = "Ayana"

	fmt.Println(myGirlName)

	var age = 24
	fmt.Println(age)

	// tipe data tanpa menulis kata kunci var
	musicGenre := "Grunge"

	fmt.Println(musicGenre)

	// jika ingin  merubah suatu variable tidak bisa menggunaka ':=', karena itu hanya berlaku untuk deklarasi awal saja
	musicGenre = "Pop Punk"
	fmt.Println(musicGenre)

	// multiple variable
	var (
		firtsName = "Muhammad"
		lastName  = "Zulfadli"
	)

	fmt.Println("First Name = ", firtsName)
	fmt.Println("Last Name = ", lastName)
}
