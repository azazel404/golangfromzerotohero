package main

import "fmt"

func main() {

	arraykedua := make([]string, 2)
	arraypertama := []string{"captain america", "captain marvel", "thor", "iron man"}
	_ = copy(arraykedua, arraypertama)

	fmt.Println(arraykedua)
	fmt.Println(arraypertama)
	// fmt.Println(dicopy)
}
