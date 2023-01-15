package main

import (
	"fmt"
)

func main() {

	AllTickets := 10

	var firstName string
	fmt.Print("What's You First Name? ")
	fmt.Scan(&firstName)		
	
	for {
		var userTickets int
		fmt.Printf("Hello %v, How many Tickets number? ", firstName)
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
