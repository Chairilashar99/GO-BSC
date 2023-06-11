package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method/ struct function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

// Cara pertama
func main() {
	var admin Customer
	admin.Name = "Heril"
	admin.Address = "Soppeng"
	admin.Age = 23

	//struct method
	admin.sayHello("Hekal")

	// fmt.Println(admin.Name)
	// fmt.Println(admin.Address)
	// fmt.Println(admin.Age)

	// //struct literal cara ke dua
	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Soppeng",
	// 	Age:     30,
	// }
	// fmt.Println(joko)

	// //Cara ketiga jarang di pake karena rentang error
	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)

}
