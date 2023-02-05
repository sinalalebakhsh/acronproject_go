//036 Slices

package main

import "fmt"

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


	for i, x := range animals  {
		fmt.Println(i, x)
	}


	//  Single Element
	fmt.Println("Element 0 is", animals[0])



	//  More than one Element
	fmt.Println("first two elements are", animals[0:2])
}
