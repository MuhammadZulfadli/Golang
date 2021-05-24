package main

import "fmt"

func main() {
	// Break example
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Ke", i)
	}

}
