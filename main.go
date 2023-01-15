//  User Input Validation
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to this Program ...[*]")
	
	// AllTickets := 10


	Counter := 0
	for Counter == 0 {
		//===============================================================================================
		var firstName string
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)		
		isValidUserName := len(firstName) < 3
		if !isValidUserName{
			Counter = 1
		} 

		//===============================================================================================
		var emailAddress string
		fmt.Print("Write Email: ")
		fmt.Scan(&emailAddress)
		isValidEmail := strings.Contains(emailAddress, "@")
		if !isValidEmail{
			Counter = 1
		}
		// //===============================================================================================
		// var userTickets int
		// fmt.Printf("Hello %v, How many Tickets number? ", firstName)
		// fmt.Scan(&userTickets)		
		// isValidTicketNumber := userTickets > 0 && AllTickets >= userTickets 
		// if !isValidTicketNumber{
		// 	Counter = 1

		// }
		// //===============================================================================================
		// var userCity string
		// fmt.Print("Write City: ")
		// fmt.Scan(&userCity)
		// isValidCity := userCity == "Tehran" || userCity == "Gilan" || userCity == "Sistan"
		// if !isValidCity{
		// 	Counter = 1

		// }
		//===============================================================================================
		

	}
}
