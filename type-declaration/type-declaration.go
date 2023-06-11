package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpHeril NoKTP = "7312060706990002"
	var marriedStatus Married = true
	fmt.Println(noKtpHeril)
	fmt.Println(marriedStatus)
}
