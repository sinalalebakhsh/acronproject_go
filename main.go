package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Introduction ()  {
	fmt.Println(`
	Guess the Number Challenge
	seed the random number generator
	Guess the Number Game
	Answer to this question:    What I need to do in this program ?
	1- Display a wellcome/instructions 
	2- take them through the games
	3- give them the answer`)
}



func main()  {
	Wellcome() 
	
	PlayTheGame()
	GoodBye()
	Introduction()
}

func Wellcome()  {
	// 1- Display a wellcome/instructions 
	fmt.Println("")
	fmt.Println("	||==========================||")
	fmt.Println("	||  Guess the Number Game   ||")
	fmt.Println("	||==========================||")
	fmt.Println("")	
}
func GoodBye()  {
	fmt.Println("")
	fmt.Println("	||==========================||")
	fmt.Println("	||         GoodBye          ||")
	fmt.Println("	||==========================||")
	fmt.Println("")	
}
func PlayTheGame() {
	const prompt = "and don't type you number here, just press ENTER when ready"
	rand.Seed(time.Now().UnixNano())
	//-------------------------------------------------------------------------
	var	firstNumber = rand.Intn(4) + 2 
	var secondNumber = rand.Intn(4) + 2
	var subtraction = rand.Intn(4) + 2
	reader := bufio.NewReader(os.Stdin)
	var answer = firstNumber * secondNumber - subtraction
	

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// 2- take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')
	
	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')
	
	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')
	
	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')
	// 3- give them the answer
	fmt.Println("The answer is", answer)

}
// Finishing Guess of the Number Game


