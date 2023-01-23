// Muliple Packages
package main

import "acronproject/helper"

// "acronproject/helper"


func main() {
	// helper.Get_Informations()
	// helper.Goodbye()

	var user_input = 30
	var user_ai = 20

	isValid := user_ai > user_input

	if isValid{
		helper.Print("Is valid")
	} else if !isValid {
		helper.Print("Is invalid")
	}


}


