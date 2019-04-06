package main

import "fmt"

//membuat object
type student struct {
	name  string
	grade int
}

func main() {
	//deklarasi object , berawal nama object yang diingin kan , spasi , nama object diatas
	var namasiswa student

	namasiswa.name = "john wick"
	namasiswa.grade = 12

	fmt.Println("name : ", namasiswa.name)
	fmt.Println("grade : ", namasiswa.grade)

}
