package main

import "fmt"

type person struct {
	alamat  string
	kodepos int
}

func main() {
	//anonymouse struct yang embed data dengan struct person
	var struct1 = struct {
		person
		grade int
	}{}

	///ngisi data struct dengan struct person
	struct1.person = person{"bengkong", 29464}
	struct1.grade = 2
	//tampilkan
	fmt.Println("alamat", struct1.person.alamat)

}
