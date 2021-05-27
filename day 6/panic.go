package main

import "fmt"

func endApp() {
	// Recover function  = Digunakan untuk menangkap data si Panic!
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	// runApp(false)
	fmt.Println("Ipay")
}
