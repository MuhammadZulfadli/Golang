package main

import "fmt"

func main() {
	// base option
	// const firtsName string = "Muhammad"
	// const lastName = "Zulfadli"
	// const value = 200

	// another option
	const (
		firtsName string = "Muhammad"
		lastName         = "Zulfadli"
		value            = 200
	)

	fmt.Println(firtsName)
	fmt.Println(lastName)
	fmt.Println(value)
}
