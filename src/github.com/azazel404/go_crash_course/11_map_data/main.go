package main

import "fmt"

func main() {

	//deklarasi map
	//cara lain
	//var chicken map[string]int
	chicken := map[string]int{}

	chicken["ayam_warna_warni"] = 50
	///jika key sama dia akan menimpa
	chicken["ayam_warna_warni"] = 100

	chicken1 := map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
		"april":    4,
		"mei":      5,
	}

	fmt.Println(chicken)
	fmt.Println(chicken1)
	delete(chicken1, "mei")
	for indexkey, val := range chicken1 {
		fmt.Println(indexkey, "     \t", val)

	}
	var value, datadicari = chicken1["april"]
	if datadicari {
		fmt.Println(value)
	} else {
		fmt.Println("data salah")
	}
}
