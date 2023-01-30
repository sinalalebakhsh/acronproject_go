package main

import(
	"fmt"
)

var one = " Is global access One"

func main()  {
	var one = "One"
	fmt.Println(one)

	// What do you think will happen ?
	MyFunc() 
}


func MyFunc()  {
	
	fmt.Println(one)
}


