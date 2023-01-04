
package main

import "fmt"

func main(){
	
	var userName string
	var userAge uint8


	fmt.Print("What is your name? ")
	fmt.Scan(&userName)

	userAge = 29
	fmt.Printf("Hello %v. with age of  %v.\n", userName, userAge)



}
