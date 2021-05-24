package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)

		counter++
	}

	// simple way
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	// Loop with slice
	slice := []string{"Muhammad", "Zulfadli", "Simatupang"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range example
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// For with map example
	person := make(map[string]string)
	person["name"] = "Zulfadli"
	person["sex"] = "Men"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
