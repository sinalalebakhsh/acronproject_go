package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready"

func main()  {
	// Guess the Number Game
	// Answer to this question:    What I need to do in this program ?
	// 1- Display a wellcome/instructions 
	// 2- take them through the games
	// 3- give them the answer
	//-------------------------------------------------------------------------
	var	firstNumber = 2 
	var secondNumber = 5
	var subtraction = 7
	var answer int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("	||==========================||")
	fmt.Println("	||  Guess the Number Game   ||")
	fmt.Println("	||==========================||")
	fmt.Println("")	
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// 2- take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')


	fmt.Println("Now multiply the result by", secondNumber, prompt)


}






