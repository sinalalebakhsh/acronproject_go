//036 Slices

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)
}
