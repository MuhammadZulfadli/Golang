package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) == true {
		fmt.Println("You're Blocked", name)
	} else {
		fmt.Println("You're Pass. Welcome!", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("ipay", blacklist)
}
