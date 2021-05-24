package main

import "fmt"

// Example for single return
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Halo " + name
	}
}

// ecample for multiple return
func getFullName() (string, string) {
	return "ipay", "Zulfadli"
}

func ignoreParameter() (string, string) {
	return "Men", "27 September 1996"
}

func main() {
	// output single return
	fmt.Println(getHello("Ipay"))

	// output multiple return
	firstName, lastname := getFullName()
	fmt.Println(firstName, lastname)

	// Outpuut ignore one of variable. Tag of ignore is "_"
	sex, _ := ignoreParameter()
	fmt.Println(sex)
}
