package main

import (
	"bufio"
	"fmt"
	"github.com/sinalalebakhsh/acronproject_go/doctor"
)

func main()  {
	
	reader := bufio.NewReader()

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}




