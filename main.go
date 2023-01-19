// Make Code cleaner with
// package level variables
// Defined at the top outside all functions
// They can be accesses inside any of the functions
//  And in all files, which are in the same package
// Belike Python language is global variable
// Local Variables:
// 		> Difined inside a function or a block
// 		> They can be accessed only inside that function or block of code
package main

import (
	"bufio"
	"fmt"
	"os"
)

var global_variable = "Sina Lalebakhsh...[*]"

func main() {
	Wellcome()
	first_name, country_name  := get_Informations()

	fmt.Printf("%v\n%v\n", first_name, country_name)
	Goodbye()
}

func Wellcome() {
	fmt.Println(" ")
	fmt.Println("Welcome to this Program ...[*]")
	fmt.Println(" ")
	fmt.Printf("Program by: %v\n", global_variable)
	fmt.Println(" ")
	fmt.Println(" ")
	
}

func get_Informations() (string, string){

	get_first_name := get_first_name()

	get_country_name := get_country_name()

	return get_first_name, get_country_name
}

func get_first_name() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter First-Name: ")
	firstName, _ := reader.ReadString('\n')	
	if len(firstName) >= 5 {
		return firstName
	}
	for {
		fmt.Print("First-Name must be more than 3 characters: ")	
		firstName, _ := reader.ReadString('\n')	
		if len(firstName) < 5 {
			continue
		}else if len(firstName) >=5{
			return firstName
		} 

	}
}

func get_country_name() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Country-Name: ")
	country_name, _ := reader.ReadString('\n')	
	if len(country_name) >= 6 {
		return country_name
	}
	for {
		fmt.Print("Country-Name must be more than 4 characters: ")	
		country_name, _ := reader.ReadString('\n')	
		if len(country_name) < 6 {
			continue
		}else if len(country_name) >= 6 {
			return country_name
		} 

	}
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
