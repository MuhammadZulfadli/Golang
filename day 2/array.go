package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Zulfadli"
	names[2] = "Simatupang"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Another way
	var angka = [2]int{
		94,
		27,
	}

	fmt.Println(angka)
	fmt.Println(angka[0])
	fmt.Println(angka[1])

	var buah = [4]string{
		"mangga", "apel", "jeruk", "kurma",
	}

	fmt.Println(buah)

	// cek panjang data dalam arrray
	fmt.Println(len(buah))
}
