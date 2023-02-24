// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package askQuestion

import (
	// "acron/con_cap"
	"bufio"
	"fmt"
	"os"
	"strings"
	"acron/con_cap"

)

func AskQuestion( ChoiceOfComputer string) string {
	
	// declaring the variable using the var keyword
	var answerUser string
	fmt.Printf("What is the capital of %v\n", ChoiceOfComputer)
	
	anotherOffers(ChoiceOfComputer)

	fmt.Print("Please Enter, after your answer: ")
	

	// scanning the input by the user
	inputReader := bufio.NewReader(os.Stdin)
	answerUser, _ = inputReader.ReadString('\n')
	
	// remove the delimiter from the string
	answerUser = strings.TrimSuffix(answerUser, "\n")
	
	
	

	return strings.ToLower(answerUser)
}



func anotherOffers( ChoiceOfComputer string) {
	
	Choice1 := con_cap.CountriesCapital[con_cap.K(ChoiceOfComputer)]


	fmt.Println(Choice1)
}