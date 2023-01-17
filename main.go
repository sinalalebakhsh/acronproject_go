// Returning Multiple values from a function	
package main

import (
	"fmt"
)

func main() {
	Wellcome()
	//===============================================================================================
	firstName, userCountry, userCity := Get_User_Input()	
	//===============================================================================================

	//===============================================================================================
	Goodbye()
}


func Wellcome()  {
	fmt.Println("Welcome to this Program ...[*]")	
}

func Get_User_Input() (string, string, string) {
	for{
		var firstName string
		var userCountry string
		var userCity string

		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)
		isValidUserName := len(firstName) >= 3

		fmt.Print("What's Your Country Name? ")
		fmt.Scan(&userCountry)		
		isValidUserCountry := len(userCountry) >= 4
		
		fmt.Print("What's Yout City Name? ")
		fmt.Scan(&userCity)
		isValidUserCity := len(userCity) >= 4

		if isValidUserName && isValidUserCountry && isValidUserCity{
				return firstName, userCountry, userCity
			}
		}		
}

// ==================================================================
// func Returning_Multiple_values(Input) (Out put Finction OR returned)
// ==================================================================
func get_user_input(firstName string, userCountry string, userCity string)  {
	fmt.Println("   ")
	fmt.Printf(" Hello %v\n  ", firstName)
	fmt.Println("   ")

	switch userCountry{
		
	}

	switch userCity {
	case "Tehran":
		fmt.Printf("%v! Your city is Capital of the Iran. \n", firstName)
	case "Gilan":
		fmt.Printf("%v! Your city is in north of the Iran. \n", firstName)
	case "New-York", "Washington":
		fmt.Printf("%v! Your city is in the United States of America. \n", firstName)
	default:
		fmt.Printf("I dont know where is it %v\n", userCity)
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










