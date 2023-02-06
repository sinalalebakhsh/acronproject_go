//038 Functions

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	z := addTwoNumber(2,3)
	fmt.Println(z)
}

func addTwoNumber(x, y int) int {
	return x + y
}
