package main

import "fmt"


var myVar = "myVar"

func main()  {
// variables
// declare a package level variable for the main
// package name myVar


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
// blockVar, and PackageVar on one line
//function in packageone
}



