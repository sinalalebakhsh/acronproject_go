// Switch
//  Statement
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to this Program ...[*]")

	//===============================================================================================
	for{
	var firstName string
	fmt.Print("What's You First Name? ")
	fmt.Scan(&firstName)		
	isValidUserName := len(firstName) >= 3
		if isValidUserName{
			fmt.Printf("Hello %v\n", firstName)
			break
		}
	}	
}
