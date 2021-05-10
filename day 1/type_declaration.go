package main

import "fmt"

func main() {
	// alias tipedata
	type NoKTP string
	type isRich bool

	var NoKtpIpay NoKTP = "11231763182631723"
	var uRich = false
	fmt.Println(NoKtpIpay)
	fmt.Println(uRich)

}
