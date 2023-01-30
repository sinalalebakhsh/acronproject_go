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

		userInput, _ := reader.ReadString('\n')

	}
}