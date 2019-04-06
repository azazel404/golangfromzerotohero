package main

import "fmt"

func main() {

	//point :=8
	//deklarasi variable dan memberikan kodisi didalam if
	if point := 8; point == 8 {
		fmt.Printf("sukses point nya adalah %d\n", point)
	} else if point < 10 {
		fmt.Println("point lebih kecil dari 10")
	} else if point > 10 {
		fmt.Println("poin lebih besar dari 10")
	} else {
		fmt.Print("point tdak diinput \n")
	}

}
