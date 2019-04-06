package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//variable yang menampung slice
	var names = []string{"john", "wick"}
	printMessage("Halo", names)

	rand.Seed(time.Now().Unix())

	randomValue := randomWithRange(2, 10)
	fmt.Println("random number", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

}

func printMessage(message string, arr []string) {
	var NameString = strings.Join(arr, " ")
	fmt.Println(message, NameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max)
	// var value = rand.Int() % (max - min + 1) + min
	return value
}
