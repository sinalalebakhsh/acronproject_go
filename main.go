// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package main

import (
	"acron/con_cap"
	"fmt"
	"acron/askQuestion"
	"acron/userAnswerDiagnosis"
)



func main() {
	// 1. Choice of Computer
	ChoiceOfComputer := con_cap.ChoiceOfComputer(con_cap.CountriesCapital)

	// 2. Ask Question from User
	AskQuestion_ := askQuestion.AskQuestion(string(ChoiceOfComputer))

	fmt.Println(AskQuestion_)
	
	// 3. Diagnosis true/false user answer  
	userAnswerDiagnosis.UserAnswer(string(ChoiceOfComputer))
	
}


/*
// declaring the variable using the var keyword
var capitalOfSouthAfrica string
fmt.Println("What is the capital of South Africa")
fmt.Print("Please enter your answer: ")
// scanning the input by the user
inputReader := bufio.NewReader(os.Stdin)
capitalOfSouthAfrica, _ = inputReader.ReadString('\n')
// remove the delimiter from the string
capitalOfSouthAfrica = strings.TrimSuffix(capitalOfSouthAfrica, "\n")

*/
