//038 Functions

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {

	mySumNumbers := sumMany(1,1,2,30)

	fmt.Println(mySumNumbers)

}

// For sum infinite numbers:
func sumMany(nums ...int) int  {
	total := 0

	for _, x := range nums  {
		total = total + x
	}

	return total
}


