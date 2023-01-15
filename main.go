//  User Input Validation
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to this Program ...[*]")

	for {
		//===============================================================================================
		var firstName string
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)		
		isValidUserName := len(firstName) >= 3
		if isValidUserName{
			break
		}
		


	}
}
