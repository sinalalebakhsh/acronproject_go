package main

import (
	"fmt"
)

func main() {

	// "For-Each" Loop
	//  Iterating over a list

var userInput string
fmt.Print("What's You Name? ")
fmt.Scan(&userInput)
fmt.Printf("Hello %v\n", userInput)
fmt.Printf("Your name have memory address: %v\n", &userInput)
	
}
