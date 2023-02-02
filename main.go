package main

import (
	"fmt"
	"os"
	// "myapp/game"
)

func main() {
	// playAgain := true

	// for playAgain {
	// 	// game.Play()
	// 	playAgain = game.GetYesOrNo("Would you like to play again (y/n)?")
	// }

	// fmt.Println("")
	// fmt.Println("Goodbye.")

	fmt.Println("Home:", os.Getenv("HOME"))
	// fmt.Println("Shell:", os.Getenv("SHELL"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok  {
		fmt.Printf("The Env var SHELL is not set")
	} else  {
		fmt.Println("Shell:", shell)
	}



}