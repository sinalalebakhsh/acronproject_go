package animation

import (
	"fmt"
	"time"
	"github.com/MasterDimmy/go-cls"
)


func ClearSceenTerminal() {
	cls.CLS()
}


func Clear_123() {
	fmt.Print("\033[1A\033[K")
}

func Animation() {
	fmt.Println()
	for j := 0; j <= 2; j ++{
		Clear_123()
		if j == 0 {
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Println()
		} else if j == 1 {
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Println()
		} else {
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(90 * time.Millisecond)
			fmt.Println()
		}
	}
}