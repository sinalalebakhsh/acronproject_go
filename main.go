// Muliple Packages
package main

import (
	"acronproject/helper"
	"fmt"
)


func main() {

	helper.Print("Salam Sina")

	helper.Get_Informations()

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
