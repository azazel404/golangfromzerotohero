package main

import "fmt"

func main() {

	average := calculcate(2, 4, 5, 7, 6, 3, 2)
	msg := fmt.Sprintf("Rata rata : %.2f", average)
	fmt.Println(msg)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 5}
	avg := calculcate(numbers...)
	msg = fmt.Sprintf("rata rata : %.2f", avg)
	fmt.Println(msg)

}
//membuat function dengan parameter tanpa batas
func calculcate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	arg := float64(total) / float64(len(numbers))
	return arg
}
