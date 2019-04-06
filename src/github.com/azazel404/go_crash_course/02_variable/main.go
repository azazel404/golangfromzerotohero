package main

import "fmt"

func main() {
	//membuat variable menggunakan tipe data dan nama variable
	var FirstName string = "john"
	var LastName string = "wick"

	fmt.Printf("Helo %s %s\n", FirstName, LastName) //sama dengan printf akan tetapi memiliki struktur berbeda
	// fmt.Println("Halo", FirstName, LastName)

	//variable tanpa menggunakan tipe data
	FirstName1 := "hendra"
	LastName1 := "gunawan"

	fmt.Printf("Hello %s %s\n", FirstName1, LastName1)

	//cara assignment variable selanjutnya
	FirstName1 = "stark"

	fmt.Printf("Assigment %s", FirstName1)
}
