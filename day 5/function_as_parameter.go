package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// another way
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kucing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ipay", spamFilter)

	// another way
	filter := spamFilter
	sayHelloWithFilter("Kucing", filter)
}
