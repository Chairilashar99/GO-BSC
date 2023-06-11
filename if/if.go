package main

import "fmt"

func main() {
	name := "Hefa"

	if name == "Heril" {
		fmt.Println("Hello Heril")
	} else if name == "Hekal" {
		fmt.Println("Hello Hekal")
	} else if name == "Hedir" {
		fmt.Println("Hello Hedir")
	} else {
		fmt.Println("Hi, kenalan oiii!")
	}

	// length := len(name)
	// if length > 5 {
	// 	fmt.Println("Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	//short Statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
