//038 Functions

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "HapHap..."
	dog.NumberOfLegs = 4
	dog.Says()
}

// For sum infinite numbers:
func sumMany(nums ...int) int  {
	total := 0

	for _, x := range nums  {
		total = total + x
	}

	return total
}


// For use struct in functions: 
type Animal struct  {
	Name string
	Sound string
	NumberOfLegs int
}
func (a *Animal) Says()  {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}
func (b *Animal) Legs()  {
	fmt("A %s have %v\n")
}