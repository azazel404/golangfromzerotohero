package main

import "fmt"

func main() {

	//array multidimensi
	// 	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	var buah = [4]string{"jeruk", "manggis", "nangka", "tomat"}              //array
	var hewan = [...]string{"kucing", "anjing", "babi", "kelelawar", "ursa"} //array tanpa definisi jumlahnnya
	var benda = []string{"handphone", "tv", "laptop", "monitor"}             // slice
	var benda1 = [2][2]string{{"handphone", "tv"}, {"laptop", "monitor"}}    // array multi dimensi

	fmt.Println(buah)
	fmt.Println(hewan[1])
	fmt.Println(benda[3])
	fmt.Println(benda1[1][0])
	tangkapslice := benda[0:3] // index awal nya untuk mulai data ditampilkan , index kedua  mau sampe brpa data dipanggil
	fmt.Println(cap(tangkapslice))

	datamulaidari1 := benda[1:3]
	// fmt.Println(cap(datamulaidari1))
	// fmt.Println(len(datamulaidari1))
	// fmt.Println(datamulaidari1)
	tambahbenda := append(datamulaidari1, "singaaaa") //nambah dalam slice array
	fmt.Println(tambahbenda)
}
