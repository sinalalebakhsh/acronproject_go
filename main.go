//040 Interface
package main
import "fmt"

func main() {
	// ask a riddle !!! یک معما بپرس
	dog := Dog {
		Name: "dog",
		Sound: "haaphaap",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

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
	Riddle(&cat)
}
/*
Interfaces type
-------------------------------
When you're defining an interface, all you do are
define the functions for an interface.
all we have to do is make sure that the types 
we define have associated functions or methods of the same name as specified in the interface.
*/
type Animal interface {
	Says() string
	HowManyLegs() int
}
func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
// Dog id the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}
func (d *Dog) Says() string {
	return d.Sound
}
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}
// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}
func (d *Cat) Says() string {
	return d.Sound
}
func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}