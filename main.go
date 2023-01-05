
package main

import "fmt"

func main(){
	
	// Infinite Loops

	var nameList [] string
	var userInput string


	
	for {
		
		fmt.Print("Name? ")
		fmt.Scan(&userInput)
		nameList = append(nameList, userInput)
		fmt.Printf("nameList: %v\n", nameList)
	}

}
