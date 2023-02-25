// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package main

import (
	"acron/con_cap"
	"acron/askQuestion"
	"acron/userAnswerDiagnosis"
	"acron/animation"
)

func main() {

	animation.Animation()

	// 1. Choice of Computer
	ChoiceOfComputer_ := con_cap.ChoiceOfComputerOrg(con_cap.CountriesCapital)

	// 2. Ask Question from User
	AskQuestion_ := askQuestion.AskQuestion(string(ChoiceOfComputer_))
	
	animation.Animation()

	// 3. Diagnosis true/false user answer
	userAnswerDiagnosis.UserAnswer(AskQuestion_, string(ChoiceOfComputer_))	
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
