package main

import "fmt"

func main() {

	//deklarasi closure / anonymouse function dan didalam function menampung tipe data slice
	var getMinMax = func(n []int) (int, int) {
		var min, max int      //deklarasi variable int
		for i, e := range n { //looping array slice
			//pengecekan logika menggunakan switch
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}
