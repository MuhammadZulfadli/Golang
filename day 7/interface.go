package main

import "fmt"

type HashName interface {
	GetName() string
}

func Sayhello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Another example
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ipay Person
	ipay.Name = " Muhammad Zulfadli"

	Sayhello(ipay)

	cat := Animal{
		Name: "Meng",
	}

	Sayhello(cat)
}
