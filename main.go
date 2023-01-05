package main

import (
	"fmt"
	"strings"
)

func main(){
	
	// "For-Each" Loop
	//  Iterating over a list

	var nameList [] string
	var userInput string

	onlyFirstNames := []string{}

	
	
	for {
		
		fmt.Print("Name? ")
		fmt.Scan(&userInput)
		nameList = append(nameList, userInput)
		fmt.Printf("nameList: %v\n", nameList)
	}
	
	for index, fullNames := range nameList{

		strings.Fields(fullNames)

	}
}
