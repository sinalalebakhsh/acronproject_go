package main

import (
	"fmt"
)

func main() {

	AllTickets := 10

	for {
		var firstName string
		var userTickets int
		
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)		
		fmt.Printf("Hello %v, How many Tickets number? ", firstName)
		fmt.Scan(&userTickets)		

		if AllTickets >= userTickets {
			AllTickets = AllTickets - userTickets
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
