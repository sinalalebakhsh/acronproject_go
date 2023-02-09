//039 Channels

package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

/* 	
reference types (pointers, silence, maps, functions, channels)
-------------------------------
interface type
-------------------------------
Go routine
-------------------------------
RESULT that's how channels work.
Now by default this channel will block.
In other words, it won't actually continue past this point until we actually push something in.
به عبارت دیگر، تا زمانی که چیزی را وارد نکنیم، عملاً از این نقطه عبور نخواهد کرد
-------------------------------
Channels are mainly used to communicate between Go routines 
to pass data from one part of your program to another part of your
that's running in the background as a go routine.
*/
var keyPressChannel chan rune

func main() {
	keyPressChannel = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func(){
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChannel <- char
	}
}


func listenForKeyPress() {
	for {
		key := <-keyPressChannel
		fmt.Println("You pressed", string(key))
	}
}