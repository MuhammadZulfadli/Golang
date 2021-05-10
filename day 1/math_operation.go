package main

import "fmt"

func main() {
	var result = 20 + 1
	fmt.Println(result)

	var a = 200
	var b = 120
	var c = a - b

	fmt.Println(c)

	var d = 90
	d += 10 // d = d + 10
	fmt.Println(d)

	var e = 10
	e++ // e + 1
	fmt.Println(e)
}
