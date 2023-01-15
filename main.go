//  User Input Validation
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to this Program ...[*]")
	
	AllTickets := 10


	Counter := 0
	for{
		if Counter == 0 {
			Counter = 1
			//===============================================================================================
			var firstName string
			fmt.Print("What's You First Name? ")
			fmt.Scan(&firstName)		
			isValidUserName := len(firstName) < 3
			//===============================================================================================
			var emailAddress string
			fmt.Print("Write Email: ")
			fmt.Scan(&emailAddress)
			isValidEmail := strings.Contains(emailAddress, "@")
			//===============================================================================================
			var userTickets int
			fmt.Printf("Hello %v, How many Tickets number? ", firstName)
			fmt.Scan(&userTickets)		
			isValidTicketNumber := userTickets > 0 && AllTickets >= userTickets 
			//===============================================================================================
			var userCity string
			fmt.Print("Write City: ")
			fmt.Scan(&userCity)
			isValidCity := userCity == "Tehran" || userCity == "Gilan" || userCity == "Sistan"
			//===============================================================================================
		}else {

		}
		
	}

}
