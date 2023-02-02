package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	fmt.Print("->  ")

	userInput, _ := reader.ReadString('\n')
}