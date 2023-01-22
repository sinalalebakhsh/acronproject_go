// Muliple Packages
package main

import (
	"acronproject/helper"
	"fmt"
)

var global_variable = "Sina Lalebakhsh...[*]"

func main() {

	helper.Print("Salam Sina")

	first_name, country_name := helper.Get_Informations()

	fmt.Printf("%v\n%v\n", first_name, country_name)
	Goodbye()
}

func Goodbye() {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")
}
