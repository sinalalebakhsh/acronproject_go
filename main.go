package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main()  {
	reader = bufio.NewReader(os.Stdin)
	userInput := readSting("What's your name?")

	age := readInt("How old are you?")

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


func readInt(a string) int  {
	fmt.Println(a)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	number, err := strconv.Atoi(userInput)
	if err != nil  {
		fmt.Println("Please enter a whole number")
	}
	
	return number
}

