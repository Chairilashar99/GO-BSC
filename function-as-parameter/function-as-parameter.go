package main

import "fmt"

type Filter func(string) string

func sayHellowithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHellowithFilter("Heril", spamFilter)
	sayHellowithFilter("Anjing", spamFilter)
}
