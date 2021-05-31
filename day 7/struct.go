package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct function/method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

// another example of struct method
func (a Customer) sayHuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var ipay Customer
	ipay.Name = "Muhammad Zulfadli"
	ipay.Age = 24
	ipay.Address = "Barcelona"

	// Output struct method
	ipay.sayHi("Ipay")

	// Output another struct method
	ipay.sayHuu()

	// fmt.Println(ipay.Name)
	// fmt.Println(ipay.Address)
	// fmt.Println(ipay.Age)

	// // struct lliterals
	// ayana := Customer{
	// 	Name:    "Ayana Shahab",
	// 	Address: "Jepang",
	// 	Age:     24,
	// }

	// fmt.Println(ayana)
	// // another struct literals (Not recommended)
	// delia := Customer{"Deliaaa", "Jepang", 24}
	// fmt.Println(delia)
}
