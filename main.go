package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sinalalebakhsh/acronproject_go/doctor"
)

func main()  {
	
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	userInput, _ := reader.ReadString('\n')

	fmt.Println(userInput)
}




