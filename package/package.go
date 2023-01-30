package packageone

import "fmt"




// declare a package level variable in the packageone
// package named PackageVar
var PackageVar = "PackageVar" 



// create an exported function in package called PrintMe
func PrintMe()  {
	fmt.Println("PrintMe() func")
}