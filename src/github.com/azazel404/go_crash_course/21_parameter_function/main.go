package main

import "fmt"
import "strings"

func main() {
	var data = []string{"hendra", "jason", "ethan"}
	var dataContainsO = filterDATA(data, func(eachdata string) bool {
		return strings.Contains(eachdata, "o")
	})
	var dataLenght5 = filterDATA(data, func(eachdata string) bool {
		return len(eachdata) == 5
	})
	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]

}

//cb adalah callback
func filterDATA(data []string, cb func(string) bool) []string {
	var result []string

	for _, each := range data {
		filtered := cb(each)
		// fmt.Println(filtered)
		if filtered {
			result = append(result, each)
		}
	}
	return result
}
