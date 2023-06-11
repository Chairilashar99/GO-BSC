package main

import "fmt"

func getFullName() (string, string) {
	return "Heril", "Ashar"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}
