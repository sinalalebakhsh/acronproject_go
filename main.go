//040 Interface

package main

import "fmt"

/*
Interfaces type
-------------------------------

*/

// Dog id the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func main() {
	// ask a riddle !!! یک معما بپرس
	dog := Dog {
		Name: "dog",
		Sound: "haaphaap",
		NumberOfLegs: 4,
	}
	Riddle(dog)
}



func Riddle(d Dog) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, d.Sound, d.NumberOfLegs)
	fmt.Println(riddle)
}
