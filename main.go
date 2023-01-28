package main

import (
	"fmt"
	"github.com/sinalalebakhsh/acronproject_go/doctor"
)

func main()  {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}




