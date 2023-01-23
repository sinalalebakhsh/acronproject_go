package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Get_Informations() (){

	get_first_name := Get_first_name()

	get_country_name := Get_country_name()

	get_age_number := Get_age_number()

	get_email_address := Get_email()

	var user_Data = make(map[string]string)
	user_Data["first_name"] = get_first_name
	user_Data["country_name"] = get_country_name
	user_Data["age_number"] = strconv.FormatUint(uint64(get_age_number), 10)
	user_Data["email_address"] = get_email_address
	var data_center = make([]map[string]string, 0)
	data_center = append(data_center, user_Data)

	fmt.Printf("Data Center: %v", data_center)
}
//----------------1-------------------------------------------------------------------------
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

	}
}
//----------------2-------------------------------------------------------------------------
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

	}
}
//----------------3-------------------------------------------------------------------------
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
	}
}
//----------------4-------------------------------------------------------------------------
func Print(userinput string){
	fmt.Println(userinput)
}
//----------------5-------------------------------------------------------------------------
func Goodbye() {
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println(" [*]----THANKS.GOODBYE----[*]  ")
	fmt.Println(" [*]----------------------[*]  ")
	fmt.Println("   ")
	fmt.Println("   ")
}
//----------------6-------------------------------------------------------------------------
func Get_email() (string){
	var user_email string
	fmt.Print("Enter Email Address: ")
	fmt.Scan(&user_email)
	return user_email
}
