package main

import "fmt"

//membuat object 1
type student struct {
	name  string
	grade int
}

//object 2 nampung data object 1
type person struct {
	alamat  string
	student //ngamsbil dri struct pertama
}

func main() {

	var namasiswa person //deklarasi object , nama object baru(namasiswa) object(person)

	//inisialisasi object
	namasiswa.name = "john wick"
	namasiswa.grade = 12
	namasiswa.alamat = "bengkong harapan 2"
	namasiswa.student.grade = 2000
	//tampilkan
	fmt.Println("name : ", namasiswa.name)
	fmt.Println("grade : ", namasiswa.grade)
	fmt.Println("nama , grade , dan alamat  : ", namasiswa.name, namasiswa.grade, namasiswa.alamat)
	fmt.Println("grade telah diubah dan akses seperti object", namasiswa.student.grade)

}
