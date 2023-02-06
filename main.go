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

	for key, value := range intMap  {
		fmt.Println(key, value)
	}
}

