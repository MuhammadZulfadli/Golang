package main

import "fmt"

func main() {
	var name1 = "ipay"
	var name2 = "ipay"

	var result bool = name1 == name2

	fmt.Println(result)

	var number1 = 200
	var number2 = 300

	fmt.Println("lebih besar ga? ", number1 > number2)
	fmt.Println("lebih kecil ga?", number1 < number2)
	fmt.Println("sama ga?", number1 != number2)
}
