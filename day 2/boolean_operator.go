package main

import "fmt"

func main() {
	var nilaiUAS = 90
	var nilaiUTS = 50

	var lulusUAS bool = nilaiUAS >= 90
	var lulusUTS bool = nilaiUTS > 40

	var lulus bool = lulusUAS && lulusUTS

	fmt.Println(lulus)

	// Simple way
	fmt.Println(nilaiUAS >= 90 && nilaiUTS >= 50)
}
