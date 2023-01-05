
package main

import "fmt"

func main(){
	
	// Slice section

	var nameList [] string

	var userInput string
	fmt.Print("name? ")
	fmt.Scan(&userInput)

	nameList = append(nameList, userInput)

	fmt.Printf("nameList: %v\n", nameList)
}
