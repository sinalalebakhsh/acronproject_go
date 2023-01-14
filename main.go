package main

import (
	"fmt"
)

func main() {

	// "For-Each" Loop
	//  Iterating over a list
	listFullName := []string{}
	for {
		var firstName string
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)
		
		var lastName string
		fmt.Print("What's You Last Name? ")
		fmt.Scan(&lastName)
		
		var sina = "sina"

		listFullName = append(listFullName, firstName + " " + lastName)
		listFullName = append(listFullName, sina)

		fmt.Printf("List Full Name: %v\n", listFullName)
		

		// listFirstNames := []string{}
	}

}
