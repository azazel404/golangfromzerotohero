package main

import "fmt"
import "strings"

type person struct {
	name  string
	grade int
}

func (s person) sayHello() {
	fmt.Println("hallo", s.name)
}

// func (s person) getNametAt(i int) string {
// 	return strings.Split(s.name, " ")[i-1]
// }
func getNametAtt(i int) string {
	var s = person{}
	return strings.Split(s.name, "/")[i-1]
}

func main() {
	var s1 = person{"hendra gunawan santos malokos inoniah zero", 21}
	s1.sayHello()

	var name = s1.getNametAtt(2)
	fmt.Println("nama panggilan: ", name)

}
