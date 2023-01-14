package main

import (
	"fmt"
)

func main() {

	// "For-Each" Loop
	//  Iterating over a list
	listFullName := []string{}
	listFirstNames := []string{}
	
	for {
		var firstName string
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)
		
		var lastName string
		fmt.Print("What's You Last Name? ")
		fmt.Scan(&lastName)
		
		listFullName = append(listFullName, firstName + " " + lastName)
		fmt.Printf("List Full Name: %v\n", listFullName)
		

		listFirstNames = append(listFirstNames, firstName)
		fmt.Printf("List Full Name: %v\n", listFirstNames)


		


	}

}
