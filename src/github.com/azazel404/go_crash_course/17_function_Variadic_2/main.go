package main

import (
	"fmt"
	"strings"
)

func main() {
	//parameter variadic bisa di kombinasikan dengan parameter biasa , sayarat parameneter
	//variadic nya diposisi akhir dari variable nya

	//array slice
	hobiku := []string{"dota2", "jualan somay", "badminton"}
	// yourHobbies("hendra", "judi", "judi2", "judi3")
	yourHobbies("hendra", hobiku...)
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, "|")
	fmt.Printf("hello my name is : %s\n", name)
	fmt.Printf("my hobbies are : %s\n", hobbiesAsString) // join digunakan hampir sama seperti explode , ubah array jdi string
}
