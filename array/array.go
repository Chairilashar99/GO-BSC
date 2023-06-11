package main

import "fmt"

func main() {

	var nameSaudara [4]string

	nameSaudara[0] = "Heril"
	nameSaudara[1] = "Hekal"
	nameSaudara[2] = "Hedir"
	nameSaudara[3] = "Haifa"

	fmt.Println(nameSaudara[0])
	fmt.Println(nameSaudara[1])
	fmt.Println(nameSaudara[2])
	fmt.Println(nameSaudara[3])

	//declare secara langsung
	var values = [3]int{
		90,
		85,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(nameSaudara))
	fmt.Println(len(values))

	var lagi [20]string

	fmt.Println(len(lagi))

}
