//033 Basic Types

package main

import "log"

// basic types (numbers, string, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Sina"
	log.Println(myString)
	myString = "Kimia"



	var myBool = true
	myBool = false

	log.Println(myBool)
}
