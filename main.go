//035 Pointers

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	x := 10
	myFirstPointer := &x
	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is to Memory Location:", myFirstPointer)

}
