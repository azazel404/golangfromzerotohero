package main

import (
	"fmt"
	"math"
)

func main() {
	//tipe variable
	var diameter float64 = 15
	var diameter1 float64 = 42
	//variable kosong nampung 2 return balikan dari function hitung luas
	area, luas := hitungLuas(diameter)
	//tampilkan data
	fmt.Println(area)
	fmt.Println(luas)

	area1, luas1 := hitungLuasLingkaran(diameter1)

	fmt.Println(area1)
	fmt.Println(luas1)
}

//func hitungluas , memiliki parameter angka dan 2 return
func hitungLuas(angka float64) (float64, float64) {

	//hitung luas
	area := math.Pow(angka/2, 2)
	//hitung keliling
	keliling := math.Pi * angka
	return area, keliling
}

//function yang bisa membalikan data itu sndiri secara tanpa declarasi variable nya
func hitungLuasLingkaran(angka float64) (area float64, keliling float64) {
	area = math.Pow(angka/2, 2)
	keliling = math.Pi * angka
	return
}
