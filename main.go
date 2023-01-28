package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sinalalebakhsh/acronproject_go/doctor"
)

func main()  {
	
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for  {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
	
		userInput = strings.Replace()


		if userInput == "quit"  {
			break
			} else  {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}


	}
}




