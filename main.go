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


}