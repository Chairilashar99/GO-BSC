package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, world!")
}

func main() {
	for i := 0; i < 6; i++ {
		sayHello()
	}
}
