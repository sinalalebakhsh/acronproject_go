
package main

import "fmt"

func main(){
	var confName = "Go Conf is here"

	const neverChangeID = 1


	fmt.Println("Welcome to Acron Project. You can ", confName, ". I like your decision")
	fmt.Println("const never change in variable, but can be change in Println:", neverChangeID+100)

	

}
