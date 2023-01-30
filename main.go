package main

import "fmt"


// declare a package level variable for the main
// package name myVar
var myVar string

func main()  {
	// variables


	// declare a block variable for the main function
	// called blockVar
	blockVar := "blockVar"
	fmt.Println(blockVar)
	// declare a package level variable in the package
	// package named PackageVar

	{
		PackageVar := "PackageVar" 
		fmt.Println(PackageVar)
	}

	// create an exported function in package called PrintMe

	// in the main function, print out the values of myVar,
	fmt.Println(myVar)
	// blockVar, and PackageVar on one line
	//function in packageone
}



