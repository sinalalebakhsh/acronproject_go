
package main

import "fmt"

func main(){
	
	var person1 = "Lalebakhsh"

	var nameList  [10]string

	nameList[0] = "Sina" + " " + person1
	nameList[9] = "Miina"

	fmt.Printf("The whole array: %v\n", nameList)

	fmt.Printf("The first value: %v\n", nameList[0])
	
	fmt.Printf("Array type: %T\n", nameList)

	fmt.Printf("Array Length: %v\n", len(nameList))

	






}
