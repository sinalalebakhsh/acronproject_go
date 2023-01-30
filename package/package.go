package packageone

import "fmt"

// declare a package level variable in the packageone
// package named PackageVar
// create an exported function in package called PrintMe



var PackageVar = "PackageVar" 

func PrintMe()  {
	fmt.Println("PrintMe() func")
}