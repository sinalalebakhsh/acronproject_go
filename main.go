// Returning Multiple values from a function
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Wellcome()
	// ===============================================================================================
	get_Informations()
	// ===============================================================================================
	// ===============================================================================================
	Goodbye()
}

func Wellcome() {
	fmt.Println("Welcome to this Program ...[*]")
}

func get_Informations() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter First-Name: ")
	firstName, err := reader.ReadString('\n')
	if err != nil {
		panic(err) // Don't forget to check and handle returned errors!
	}

	for len(firstName) < 5 {
		fmt.Print("First-Name must be more than 3 characters: ")
		firstName, err := reader.ReadString('\n')
		if err != nil {
			panic(err) // Don't forget to check and handle returned errors!
		}
		if len(firstName) > 4 {
			fmt.Scan(&firstName) 
			break
		}
	}

	fmt.Println("Hello", firstName)

	// var firstName string
	// fmt.Print("What's Your First Name? ")
	// fmt.Scan(&firstName)

	// if firstName == "\n"{
	// 	for {
	// 		fmt.Print("Pleas write name more than 3 caracters, What's Your First Name? ")
	// 		fmt.Scan(&firstName)
	// 		if len(firstName) > 3 {
	// 			break
	// 		}
	// 	}
	// }

	// fmt.Printf("Your first name = %v\n", firstName)

}

func Goodbye() {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")
}
