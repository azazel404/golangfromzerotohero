package main

import "fmt"

func main() {
	number1 := 10

	switch {
	case number1 == 8:
		fmt.Println("perfect")
	case (number1 < 8) && (number1 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	number2 := 11
	switch number2 {
	case 10:
		fmt.Println("salah nomornya yang dimasukin 10")
	case 11:
		fmt.Println("sukses ini benar gan")
	default:
		fmt.Println("failed led led")
	}
}
