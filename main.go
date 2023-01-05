package main

import (
	"fmt"
	"strings"
)

func main() {

	// "For-Each" Loop
	//  Iterating over a list

	var nameList []string
	var userInput string

	onlyFirstNames := []string{}


	for _, fullNames := range nameList {

		var getFullName = strings.Fields(fullNames)
		onlyFirstNames = append(onlyFirstNames, getFullName[0])

	}
}
