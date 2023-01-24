// Muliple Packages
package main

import (
	"acronproject/helper"
	"fmt"
)

// "acronproject/helper"


func main() {

	helper.Get_Informations()
	helper.Goodbye()

	type user_data struct {
		first_name string
		country_name string
		age_numner string 
	}

	var data_center = make([]user_data, 0)
	fmt.Println(data_center)

}


