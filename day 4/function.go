package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

// function with parameter
func sayHelloTo(firstName string, lastname string) {
	fmt.Println("Halo", firstName, lastname)
}

// another example of function parameter
func oddEven(angka int) {
	for i := 1; i <= angka; i++ {
		if i%2 == 0 {
			fmt.Println(i, "adalah bilang genap")
		} else {
			fmt.Println(i, "adalah bilanga ganjil")
		}
	}
}

func fizzBuzz(total int) {
	for i := 1; i <= total; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func main() {
	sayHello()
	sayHelloTo("Muhammad", "Zulfadli")
	oddEven(10)
	fizzBuzz(100)
}
