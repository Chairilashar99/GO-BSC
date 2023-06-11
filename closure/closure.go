package main

import "fmt"

func main() {
	name := "Heril"
	counter := 0

	//closure = scope yang di atas dapat mengakses scope yang di dalam sedangkan scope yang di dalam tidak dapat mengakses scope yang ada di luar
	increment := func() {
		name := "Hekal"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
