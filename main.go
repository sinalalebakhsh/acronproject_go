package main

import (
	"fmt"
	"log"

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

	fmt.Println("Press ESC to quit.")
	fmt.Println("MENU")
	fmt.Println("====")
	fmt.Println("1- Cappucino")
	fmt.Println("2- Dark Coffee")
	fmt.Println("3- Americano")

	for  {
		char, key, err := keyboard.GetSingleKey()
		if err != nil  {
			log.Fatal(err)
		}

		if key != 0  {
			fmt.Println("You pressed", char, key)
		} else  {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc  {
			break
		}
	}


	fmt.Println(" Program exiting ... ")

}