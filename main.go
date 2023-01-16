// Returning Multiple values from a function	
package main

import (
	"fmt"
)

func main() {
	Wellcome()
	//===============================================================================================
	Get_User_Input()	
	//===============================================================================================

	//===============================================================================================
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
// ==================================================================
// func Returning_Multiple_values(Input) (Out put Finction OR returned)
// ==================================================================
func Returning_Multiple_values() (string, string, string){
	firstValue := "firstValue"
	secondValue := "secondValue" 
	thirdValue := "thirdValue"
	return firstValue, secondValue, thirdValue
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










