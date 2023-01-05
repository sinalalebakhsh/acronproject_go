
package main

import "fmt"

func main(){
	
	// Loops

	var nameList [] string
	var userInput string


	fmt.Print("Name? ")
	fmt.Scan(&userInput)
	nameList = append(nameList, userInput)
	fmt.Printf("nameList: %v\n", nameList)

}
