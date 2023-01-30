package main

import (
	"bufio"
	"fmt"
	"os"
)



func main()  {
	reader := bufio.NewReader(os.Stdin)

	for  {
		fmt.Print("-> ")

		userInput, sina := reader.ReadString('\n')

		fmt.Print(userInput,  sina)
	}
}