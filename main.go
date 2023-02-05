//034 Aggregate Types

package main

import "log"

// basic types (numbers, string, booleans)

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxery        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {

	var myStrings [3]int
	myStrings[0] = 0
	myStrings[1] = 1
	myStrings[2] = 2
	log.Println("First Elements is:", myStrings[0])
}
