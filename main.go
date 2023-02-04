// Challenge 
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/eiannone/keyboard"
	"log"
)

var reader *bufio.Reader
type User struct  {
	UserName 		string
	Age 			int
	FavouriteNumber float64
	OwnADog 		bool
}

func main()  {
	playAgain := true

	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readSting("What's your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnADog = OwnADog("Do you have a dog?")
	
	fmt.Println(fmt.Sprintf("Your name's %s, and you'r %d years old. Your Favourite number's %.2f. Owns a dog: %t", 
	user.UserName, 
	user.Age,
	user.FavouriteNumber,
	playAgain,
	))
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
	
		number, err := strconv.ParseFloat(userInput, 64)
		if err != nil  {
			fmt.Println("Please enter a number")
		} else  {
			return number
		}
	}
}

func OwnADog(q string) bool  {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" &&  strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else  if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y'  {
			return true
		}
	}
}
