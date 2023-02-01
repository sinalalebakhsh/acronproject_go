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


	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Dark"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Latte"
	coffees[7] = "Espresso"

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
		t := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t)

		if char == 'q' || char == 'Q'  {
			break
		}
	}


	fmt.Println(" Program exiting ... ")

}