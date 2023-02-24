package userAnswerDiagnosis

import (
	"acron/con_cap"
	"fmt"
)

func UserAnswer(ChoiceOfComputer string)  {
	GetTrueAnswer :=   con_cap.CountriesCapital[con_cap.K(ChoiceOfComputer)]
	
	fmt.Println("capital of", ChoiceOfComputer, "is", GetTrueAnswer)
}