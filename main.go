//034 Aggregate Types

package main

import "log"

// basic types (numbers, string, booleans)

// aggregate types (array, struct)

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {

	var myStrings [3]int
	myStrings[0] = "cat1"
	myStrings[1] = "dog2"
	myStrings[2] = "fish3"


	log.Println("First Elements is:", myStrings[0])
}
