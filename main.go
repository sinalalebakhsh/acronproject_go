// Expermenting with String Interpolation 027
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader
type User struct  {
	UserName string
	Age int
	FavouriteNumber float64
}

func main()  {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.UserName = readSting("What's your name?")
	user.Age = readInt("How old are you?")

	fmt.Println(fmt.Sprintf("Your name's %s, and you'r %d years old.", user.UserName, user.Age))
}

func prompt()  {
	fmt.Print("->  ")
}

func readSting(a string) string  {	
	for  {
		fmt.Println(a)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput != ""  {
			return userInput
		} else  {
			fmt.Println("Please enter name with a value")
		}
	}
}


func readInt(a string) int  {
	for  {
		fmt.Println(a)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
	
		number, err := strconv.Atoi(userInput)
		if err != nil  {
			fmt.Println("Please enter a whole number")
		} else  {
			return number
		}
	}
	
}

func readFloat(a string) float64  {
	for  {
		fmt.Println(a)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
	
		number, err := strconv.ParseFloat(userInput)
		if err != nil  {
			fmt.Println("Please enter a whole number")
		} else  {
			return number
		}
	}
	
}

