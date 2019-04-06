package main

import "fmt"

func main() {

	var numberA int = 4
	fmt.Println("before", numberA)

	change(&numberA, 10)
	fmt.Println("After", numberA)

}
func change(original *int, value int) {
	*original = value
}
