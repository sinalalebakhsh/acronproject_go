package helper

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Get_Informations() (){

	TimeDelay3()	
	get_first_name := Get_first_name()	
	
	TimeDelay3()
	get_country_name := Get_country_name()

	TimeDelay3()
	get_age_number := Get_age_number()
	
	type User_Data struct {
		first_name string
		country_name string
		age_numner int
	}

	var user_Data = User_Data{
		first_name: get_first_name,
		country_name: get_country_name,
		age_numner: get_age_number,
	}

	
	TimeDelay3()
	fmt.Printf("Data Center:\n{First name: %vCountry name: %vAge number: %v\n}", 
		user_Data.first_name, 
		user_Data.country_name,
		user_Data.age_numner)

	TimeDelay3()
	fmt.Println("")}
	

func Get_first_name() (string) {
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

	}}

func Get_country_name() (string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Country-Name: ")
	country_name, _ := reader.ReadString('\n')	
	if len(country_name) >= 5 {
		return country_name
	}
	for {
		fmt.Print("Country-Name must be more than 4 characters: ")	
		country_name, _ := reader.ReadString('\n')	
		if len(country_name) < 5 {
			continue
		}else if len(country_name) >= 5 {
			return country_name
		} 

	}}

func Get_age_number() (int) {
	var age_number int
	fmt.Print("Enter Age number: ")
	fmt.Scan(&age_number)	
	if age_number >= 10 {
		return age_number
	}
	for {
		fmt.Print("Age Number must be more than 10 age: \n")
		if age_number < 10 {
			continue
		}else if age_number >= 10 {
			return age_number
		} 
	}}

func Print(userinput string){
	fmt.Println(userinput)}

func Goodbye() {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")}

func TimeDelay3()  {
	fmt.Println(".")
	time.Sleep(1 * time.Second)
	fmt.Println(".")
	time.Sleep(1 * time.Second)
	fmt.Println(".")
	time.Sleep(1 * time.Second)
}