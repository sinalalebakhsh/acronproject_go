//038 Functions

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	z := addTwoNumber(2,3)
	fmt.Println(z)
}

// func addTwoNumber(x, y int) int {
// 	return x + y
// }

//Also you can write like it:
// But it's bad readable for another programmer
func addTwoNumber(x, y int) (sum int) {
	sum = x + y
	return
}





