// Returning Multiple values from a function
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Wellcome()
	sina := get_Informations()
	fmt.Printf("%v", sina)
	Goodbye()
}

func Wellcome() {
	fmt.Println("Welcome to this Program ...[*]")
}

func get_Informations() (string){
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

func Goodbye() {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")
}
