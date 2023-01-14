package main

import (
	"fmt"
)

func main() {

	// "For-Each" Loop
	//  Iterating over a list

var firstName string
fmt.Print("What's You First Name? ")
fmt.Scan(&firstName)

var lastName string
fmt.Print("What's You Last Name? ")
fmt.Scan(&lastName)


fmt.Printf("Hello %v %v\n", firstName, lastName)
fmt.Printf("First name memory address: %v\n", &firstName)
fmt.Printf("Last name memory address: %v\n", &lastName)
	
}
