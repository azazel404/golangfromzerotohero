package main

import "fmt"

func main() {

	var arraylarik = [...]int{1, 2, 3, 4, 5, 66, 3} // array 1 dimensi
	fmt.Println(len(arraylarik))                    // ngitung array

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}} // array multidimensi
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}             // array multi dimensi
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//penggunaa for range
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}


	//error tidak bisa di append dalam array
	// names := [...]string{"trafalgard", "d", "water", "law"}
	// saya := append(names, "papayadong gan")
	// fmt.Println(saya)
}
