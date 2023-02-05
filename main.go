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


	for _, x := range animals  {
		fmt.Println(x)
	}
}
