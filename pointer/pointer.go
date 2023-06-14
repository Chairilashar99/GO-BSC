package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// function pointer
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//Pass by Value maksdnya mengcopy data pertama tanpa harus mengubah datanya
	// address2 := address1

	//Pass By reference maksudnya merubah ke dua datanya
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	*address2 = Address{
		"Makassar", "Sulawesi", "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// var address4 *Address = new(Address)
	// fmt.Println(address4)

	address4 := new(Address)
	address4.City = "Soppeng"
	fmt.Println(address4)

	//function pointer
	var alamat = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
