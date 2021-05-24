package main

import "fmt"

func getFullNamme() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Zulfadli"
	lastName = "Simatupang"

	return
}

func main() {
	namaDepan, namaTengah, namaBelakang := getFullNamme()

	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(namaBelakang)

	fmt.Println(getFullNamme())
}
