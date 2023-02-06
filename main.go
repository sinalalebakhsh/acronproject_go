//037 Maps

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {

	// var myMap map[string]string

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// In the Go language, reading by the loop is done randomly.
	// This behavior is weird !!!
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	fmt.Println("--------------")
	// For delete an element in Map
	// If comment it below first code
	// delete(intMap, "five")

	element, ok := intMap["five"]
	if ok {
		fmt.Println(element, "exsit In intMap.")
	} else {
		fmt.Println(element, "exsit not In intMap.")
	}

	// For change value of an index Map
	intMap["five"] = 333
	element2, ok := intMap["five"]
	if ok {
		fmt.Println(element2, "exsit In intMap.")
	} else {
		fmt.Println(element2, "exsit not In intMap.")
	}

}
