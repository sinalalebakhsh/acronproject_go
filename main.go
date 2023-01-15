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
			var firstName string
			fmt.Print("What's You First Name? ")
			fmt.Scan(&firstName)		
		
			isValidUserName := len(userTickets) < 3
		
			var emailAddress string
			fmt.Print("Write Email: ")
			fmt.Scan(&emailAddress)
			isValidEmail := strings.Contains(emailAddress, "@")
			
			Counter = 1
			var userTickets int
			fmt.Printf("Hello %v, How many Tickets number? ", firstName)
			fmt.Scan(&userTickets)		
			
			isValidTicketNumber := userTickets > 0 && AllTickets >= userTickets 

			if AllTickets > userTickets {
				AllTickets = AllTickets - userTickets
			}else if AllTickets == userTickets {
				fmt.Printf("Oh!! You get every tickets! Thank for buying ... [*]\n")
				break
			}else{
				fmt.Print("Teckets Number more than All teckets ...[*] \n")
				break
			}
				
			if AllTickets == 0 {
				// End Loop:
				fmt.Printf("You Get %v Teckets, So is Finished...[*]\n ", userTickets)
				break
			}
			}else {
			var userTickets int
			fmt.Print("How many Tickets number? ")
			fmt.Scan(&userTickets)		
	
			if AllTickets > userTickets {
				AllTickets = AllTickets - userTickets
			}else if AllTickets == userTickets {
				fmt.Printf("Oh!! You get every tickets! Thank for buying ... [*]\n")
				break
			}else{
				fmt.Print("Teckets Number more than All teckets ...[*] \n")
				break
			}
	
			if AllTickets == 0 {
				// End Loop:
				fmt.Printf("You Get %v Teckets, So is Finished...[*]\n ", userTickets)
				break
			}
		}
		
	}

}
