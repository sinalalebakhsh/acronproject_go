// Returning values from a function	
package main

import (
	"fmt"
)

func main() {
	Wellcome()
	//===============================================================================================
	Get_User_Input()	

	user_input := amount_of_money()

	result_amount_of_money(user_input)

	Goodbye()
}


func Wellcome()  {
	fmt.Println("Welcome to this Program ...[*]")	
}

func Get_User_Input()  {
	
	for{
		var firstName string
		var userCity string
	
		fmt.Print("What's You First Name? ")
		fmt.Scan(&firstName)	
		isValidUserName := len(firstName) >= 3
		if isValidUserName{
				fmt.Print("What's Your City Name? ")
				fmt.Scan(&userCity)		
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
				break
			}
		}
}

func amount_of_money() int {
	var user_input int
	fmt.Print("How much dollers with you? ")
	fmt.Scan(&user_input)
	return user_input
}

func result_amount_of_money(user_input int){
	if user_input < 10{
		fmt.Printf("%v dollers is not enough... [*] \n", user_input)
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
