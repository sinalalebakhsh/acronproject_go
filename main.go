// Returning Multiple values from a function	
package main

import (
	"fmt"
)

func main() {
	Wellcome()	
	// ===============================================================================================
	get_user_Input()	
	// ===============================================================================================	
	// ===============================================================================================
	Goodbye()
}


func Wellcome()  {
	fmt.Println("Welcome to this Program ...[*]")	
}

func get_user_Input()  {

	var firstName string
	
	fmt.Print("What's You First Name? ")
	fmt.Scan(&firstName)

	for len(firstName) <= 3 {
		fmt.Print("Pleas write name more than 3 caracters, What's You First Name? ")
		fmt.Scan(&firstName)	
	}
	
}

func Goodbye()  {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")
}










