// Make Code cleaner with
// package level variables
// Defined at the top outside all functions
// They can be accesses inside any of the functions
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Wellcome()
	first_name, country_name  := get_Informations()

	fmt.Printf("%v\n%v\n", first_name, country_name)
	Goodbye()
}

func Wellcome() {
	fmt.Println("Welcome to this Program ...[*]")
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
