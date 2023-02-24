// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package askQuestion

import (
	"fmt"
	"bufio"
	"os"
	"strings"

)

func AskQuestion( ChoiceOfComputer string) string {
	
	// declaring the variable using the var keyword
	var answerUser string
	fmt.Printf("What is the capital of %v\n", ChoiceOfComputer)
	fmt.Print("Please Enter, after your answer: ")
	
	// scanning the input by the user
	inputReader := bufio.NewReader(os.Stdin)
	answerUser, _ = inputReader.ReadString('\n')
	
	// remove the delimiter from the string
	answerUser = strings.TrimSuffix(answerUser, "\n")
	
	return strings.ToLower(answerUser)
}

