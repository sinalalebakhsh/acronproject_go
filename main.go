package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	fmt.Print("->  ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	fmt.Println("Your name is", userInput+ ".")
}