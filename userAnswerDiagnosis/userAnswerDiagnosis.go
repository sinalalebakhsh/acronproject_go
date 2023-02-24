package userAnswerDiagnosis

import (
	"acron/con_cap"
	"fmt"
)

func UserAnswer(ChoiceOfComputer string) con_cap.V {
	GetTrueAnswer :=   con_cap.CountriesCapital[con_cap.K(ChoiceOfComputer)]
	
	fmt.Println("capital of", ChoiceOfComputer, "is", GetTrueAnswer)
	fmt.Printf("Type of, GetTrueAnswer is %T\n", GetTrueAnswer )
	
	return GetTrueAnswer
}