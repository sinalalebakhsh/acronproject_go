//034 Aggregate Types

package main

import "fmt"

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
	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.Luxery = true
	// myCar.Make = "mercedes-benz-unimog"

	// ---------------------------------------- Another Definition
	myCar1 := Car{
		NumberOfTires: 4,
		Luxery:        true,
		BucketSeats:   true,
		Make:          "mercedes-benz-unimog",
		Year:          2030,
	}

	fmt.Println(myCar1)
}
