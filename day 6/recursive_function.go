package main

import "fmt"

// Faktorial loop
func faktorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i //result dikalikan dengan i
	}
	return result
}

// recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	// output faktorial loop (with no recursive)
	loop := faktorialLoop(5)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	// output with recursive function
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
