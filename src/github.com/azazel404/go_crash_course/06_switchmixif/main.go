package main

import "fmt"

func main() {

	point := 0

	if point > 7 { //jika point lebih besar dari 7
		switch point {
		case 10: // dan point nya itu sebesar 10 , masuk tampilkan true validasi
			fmt.Println("sukses masuk validasi ke 7")
		default:
			fmt.Println("gagal validasi di 10")
		}
	} else { //jika gagal
		if point == 5 { // point nilai nya 5
			fmt.Println("sukses validasi 5")
		} else if point == 3 { //point nilai nya 3
			fmt.Println("sukses validasi di 3")
		} else { // gagal
			fmt.Println("you can do it") //tampilan else
			if point == 0 {              //jika sudha salah dan nilai nya 0 , tampilkan println
				fmt.Println("try harder!")
			}
		}
	}
}
