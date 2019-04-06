package main

import "fmt"

func main() {

	//tipe data interger
	var PositifNumber uint64 = 89
	var negativeNumber = -1243423644

	fmt.Println(PositifNumber)  //bilang positif
	fmt.Println(negativeNumber) //bilangan negatif

	//tipe data decimal
	var decimalNumber = 2.62
	// %nf mengubah float jdi string dan mengubah beberapa angka yng ditampilkan
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
}
