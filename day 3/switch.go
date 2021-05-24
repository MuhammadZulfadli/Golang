package main

import "fmt"

func main() {
	// Default switch statement
	name := "ipay"

	switch name {
	case "ipay":
		fmt.Println("Halo ipay")
	case "zulfadli":
		fmt.Println("Halo Zulfadli")
	default:
		fmt.Println("Kenalan lah dulu, sombong kali pun!")
	}

	// another way with short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
