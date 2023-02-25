package userAnswerDiagnosis

import (
	"acron/con_cap"
	"fmt"
)

func UserAnswer(AskQuestion_, ChoiceOfComputer string) {
	GetTrueAnswer := con_cap.CountriesCapital[con_cap.K(ChoiceOfComputer)]
	
	if string(GetTrueAnswer) == AskQuestion_ {
		fmt.Println("Yes! Capital of", ChoiceOfComputer, "is", GetTrueAnswer + ".")
	} else if string(GetTrueAnswer) != AskQuestion_ {
		fmt.Println("No! Capital of", ChoiceOfComputer, "is'n", AskQuestion_ + ".\nBut is", GetTrueAnswer + ".")
	}
}