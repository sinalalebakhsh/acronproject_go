package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	fmt.Println("Press any key on the keyboard. Or press ESC to quit. ")

	for  {
		char, key, err := keyboard.GetSingleKey()
		if err != nil  {
			log.Fatal(err)
		}

		if key != 0  {
			
		} else  {

		}
	}




}