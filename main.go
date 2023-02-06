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

	//------------------- another way for get from struct
	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true
	/*
	I can't pass this type to this function because
	this expects to see a dog and this is where interfaces come in and they become very useful.
	*/
	Riddle(cat)
}



func Riddle(d Dog) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, d.Sound, d.NumberOfLegs)
	fmt.Println(riddle)
}
