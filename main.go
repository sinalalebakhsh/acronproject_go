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
	//--------------------------------------------------------------------------------------------------
	
	const prompt = "and don't type you number here, just press ENTER when ready"
	rand.Seed(time.Now().UnixNano())
	//-------------------------------------------------------------------------
	var	firstNumber = rand.Intn(4) + 2 
	var secondNumber = rand.Intn(4) + 2
	var subtraction = rand.Intn(4) + 2
	reader := bufio.NewReader(os.Stdin)
	var answer = firstNumber * secondNumber - subtraction
	


	//--------------------------------------------------------------------------------------------------
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
func PlayTheGame(firstNumber, secondNumber, subtraction, answer int) {

}
// Finishing Guess of the Number Game


