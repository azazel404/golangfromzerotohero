package main

import "fmt"

func main() {

	chicken := []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "female"},
		map[string]string{"name": "chicken black", "gender": "male"},
	}

	for _, chickenrepeat := range chicken {
		fmt.Println(chickenrepeat["gender"], chickenrepeat["name"])
	}
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, datarepeat := range data {
		fmt.Println(datarepeat)
	}
}
