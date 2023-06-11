package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 15
	var b = 15
	var c = a * b
	fmt.Println(c)

	//Augmented Assignments
	var i = 25
	i += 25 // i = i + 25
	fmt.Println(i)

	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = 100
	fmt.Println(positive)
	fmt.Println(negative)
}
