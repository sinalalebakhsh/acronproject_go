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

	for  {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
	
		response := doctor.Response(userInput)

		fmt.Println(response)
	}
}




