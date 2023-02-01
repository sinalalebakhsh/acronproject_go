package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main()  {
	err := keyboard.Open()
	if err != nil  {
		log.Fatal(err)
	}

	defer func ()  {
		_ = keyboard.Close()
	}()


	coffees := make(map)

	fmt.Println("MENU")
	fmt.Println("====")
	fmt.Println("1- Cappucino")
	fmt.Println("2- Dark")
	fmt.Println("3- Americano")
	fmt.Println("4- Mocha")
	fmt.Println("5- Macchiato")
	fmt.Println("6- Latte")
	fmt.Println("7- Espresso")
	fmt.Println("Q- Quit the program")

	for  {
		char, _, err := keyboard.GetSingleKey()
		if err != nil  {
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(string(char))

		t := fmt.Sprintf("You chose %d", i)
		fmt.Println(t)

		if char == 'q' || char == 'Q'  {
			break
		}
	}


	fmt.Println(" Program exiting ... ")

}