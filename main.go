package main

import (
	"fmt"
	"acronproject/package"
)

// declare a package level variable for the main
// package name myVar
var myVar = "Package level variable"

func main() {
	// variables

	// declare a block variable for the main function
	// called blockVar


	// in the main function, print out the values of myVar,
	fmt.Println(myVar)
	// blockVar, and PackageVar on one line
	//function in packageone
	packageone.PrintMe()
}
