package main

import "fmt"

func main() {

	name := "Heri"

	switch name {
	case "Heril":
		fmt.Println("Hello Heril")
		fmt.Println("Hello Heril")
	case "Hekal":
		fmt.Println("Hello Hekal")
		fmt.Println("Hello Hekal")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//short statement di switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//mirip kondidsi if
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
