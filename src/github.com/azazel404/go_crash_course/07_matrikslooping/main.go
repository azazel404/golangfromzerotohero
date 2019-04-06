package main

import "fmt"

func main() {

outerLoop: //label
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]\n")
		}
	}

}
