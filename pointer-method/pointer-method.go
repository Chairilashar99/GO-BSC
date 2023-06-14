package main

import "fmt"

// pada saat menggunakan struct method atau struct function usahakan menggunakan pointer
// kelebihannnya lebih hemat memori karena tdk menduplicat" datanya
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	heril := Man{"Heril"}
	heril.Married()

	fmt.Println(heril.Name)
}
