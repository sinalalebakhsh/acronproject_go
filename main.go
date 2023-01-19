// Returning Multiple values from a function	
package main

import (
	"fmt"
)

func main() {
	Wellcome()	
	// ===============================================================================================
	get_Informations()	
	// ===============================================================================================	
	// ===============================================================================================
	Goodbye()
}


func Wellcome()  {
	fmt.Println("Welcome to this Program ...[*]")	
}

func get_Informations()  {

	var firstName string
	fmt.Print("What's You First Name? ")
	fmt.Scan(&firstName)
	for {
		if len(firstName) <= 3{
			fmt.Print("Pleas write name more than 3 caracters, What's You First Name? ")
			fmt.Scan(&firstName)	
		}else if len(firstName) > 3 {
			break
		}
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










