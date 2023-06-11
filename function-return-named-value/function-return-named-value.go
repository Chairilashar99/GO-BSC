package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "first"
	middleName = "middle"
	lastName = "last"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
