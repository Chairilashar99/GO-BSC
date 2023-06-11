package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//short statement
	for counter := 1; counter <= 11; counter++ {
		fmt.Println("Perulangan ke", counter)

	}

	slice := []string{"Heril", "Hekal", "Hedir", "Haifa"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range bisa di gunakan pada slice, array, map
	// for i, value := range slice {
	// 	fmt.Println("Indexnya", i, "=", value)
	// }

	//for range slice
	for _, value := range slice {
		//jika ada satu variable yang nda di pakai bisa menggunakan underscores _,
		// fmt.Println("Indexnya", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Heril"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
