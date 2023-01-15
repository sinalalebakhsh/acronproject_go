package main

import (
	"fmt"
)

func main() {

	AllTickets := 10

	for {
		var firstName string
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)		

		var userTickets int
		fmt.Print("How many Tickets number? ")
		fmt.Scan(&userTickets)		
		
		if AllTickets < userTickets {
			fmt.Print("user Number more than All...[*] \n")
			break
		}
		
		AllTickets = AllTickets - userTickets


		if AllTickets == 0 {
			// End Loop:
			fmt.Print("Loop's Ended...[*]\n ")
			break
		}
		
	}
}
