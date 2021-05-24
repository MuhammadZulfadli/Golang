package main

import "fmt"

func main() {
	// If else statement
	name := "Zulfadli"
	if name == "Ipay" {
		fmt.Println("Halo Ipay")
	} else {
		fmt.Println("Nama Kau ipay, gausah banyak tingkah!")
	}

	// If else if
	nilai := 90
	if nilai == 50 {
		fmt.Println("Nilai kamu C, Belum lulus!")
	} else if nilai == 70 {
		fmt.Println("Nilai kamu B, Lulus dong..")
	} else {
		fmt.Println("Nilai kamu A. Wah bagus, ga nyontek kan?")
	}

	// If Short statement
	var length = len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// another way, short statement code
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	}
}
