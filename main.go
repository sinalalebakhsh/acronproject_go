package main

import (
	"fmt"
	"strings"
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
		
		listFullName = append(listFullName, firstName + " " + lastName)
		fmt.Printf("List Full Name: %v\n", listFullName)
		
		
		// First Way
		// listFirstNames = append(listFirstNames, firstName)
		// fmt.Printf("List Full Name: %v\n", listFirstNames)
		
		// Second Way
		listFirstNames := []string{}
		for _, FullName := range listFullName{
			var FullName_ = strings.Fields(FullName)
			listFirstNames = append(listFirstNames, FullName_[0])
		}

		fmt.Printf("list First Names: %v\n", listFirstNames)
		


	}

}
