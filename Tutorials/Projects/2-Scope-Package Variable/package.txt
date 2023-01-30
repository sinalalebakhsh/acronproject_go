package packageone

import "fmt"

// declare a package level variable in the packageone
// package named PackageVar
var PackageVar = "PackageVar" 



// create an exported function in package called PrintMe
func PrintMe(s1, s2 string)  {
	fmt.Println(s1, s2, PackageVar)
}