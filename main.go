package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main()  {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	prompt()
	userInput, _ := reader.ReadString('\n')

	fmt.Println("Your name is", userInput + ".")
}

func prompt()  {
	fmt.Print("->  ")
}

func readSting(a string) string  {	

	fmt.Println(a)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
