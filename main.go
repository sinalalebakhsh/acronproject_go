
package main

import "fmt"

func main(){
	
	var person1 = "Lalebakhsh"

	var nameList  [5]string

	nameList[0] = "Sina" + " " + person1
	nameList[4] = "Miina"

	fmt.Printf("The whole array: %v\n", nameList)

	fmt.Printf("The first value: %v\n", nameList[0])
	
	fmt.Printf("Array type: %T\n", nameList)

	fmt.Printf("Array Length: %v\n", len(nameList))

	






}
