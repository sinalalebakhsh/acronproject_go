
package main

import "fmt"

func main(){
	
	var firstName string
	var lastName string
	var emailAddress string
	var phoneNumber string

	fmt.Print("First Name? ")
	fmt.Scan(&firstName)
	fmt.Print("Last Name? ")
	fmt.Scan(&lastName)
	fmt.Print("Email Address? ")
	fmt.Scan(&emailAddress)
	fmt.Print("Phone Number? ")
	fmt.Scan(&phoneNumber)

	fmt.Printf("Hello %v %v. Your Email is %v, and Phone number is %v\n", firstName, lastName, emailAddress, phoneNumber)



}
