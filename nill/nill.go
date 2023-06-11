package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// func main() {
// 	person := NewMap("Heril")
// 	fmt.Println(person)
// }

func main() {
	var person map[string]string = NewMap("Heril")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
