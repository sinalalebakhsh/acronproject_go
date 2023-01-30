package main

import(
	"fmt"
)

func main()  {
	var one = "One"
	
	fmt.Println(one)

	// What do you think will happen ?
	MyFunc() 
}


func MyFunc()  {
	var one = "the number one"
	fmt.Println(one)
}


