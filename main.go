//036 Slices

package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {

	// type MyStruct struct  {
	// 	name string
	// 	sureName string
	// }
	// fullName := MyStruct {
	// 	name: "Sina",
	// 	sureName: "Lalebakhsh",
	// }

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	//  Single Element
	fmt.Println("Element 0 is", animals[0])

	//  More than one Element
	fmt.Println("first two elements are", animals[0:2])

	// How many elements we have ?
	fmt.Println("The slice is", len(animals), "elements long.")

	// Diognosis sort a slice
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	// sorting a slice
	sort.Strings(animals)

	// showing all elements in slice
	fmt.Println("first two elements are", animals[:])

	// Diognosis sort a slice again
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	// Delete a element from slice
	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
