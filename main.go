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
	var myCar Car
	myCar.NumberOfTires = 4
	myCar.Luxery = false
	myCar.Make = "mercedes-benz-unimog"
}
