package main

import "fmt"

type student struct {
	nama  string
	grade string
}

func main() {
	var s1 = student{} //deklarasi object kosong
	//inisialisasi object kosong
	s1.nama = "hendra"
	s1.grade = "S1"

	//inisialisai object kosong cara new
	var s2 = student{"ethan", "S2"}
	// var _ = student{nama: "jason"}

	fmt.Println("student 1 :", s1.nama)
	fmt.Println("student 2 :", s2.nama)
}
