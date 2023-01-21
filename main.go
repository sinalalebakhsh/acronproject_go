// Muliple Packages
package main

import (
	"fmt"
)

var global_variable = "Sina Lalebakhsh...[*]"

func main() {

	print("Salam Sina")



	first_name, country_name  := get_Informations()

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
