package animation

import (
	"fmt"
	"time"
)


func Clear_123() {
	fmt.Print("\033[1A\033[K")
}

func Animation() {
	fmt.Println()
	for j := 0; j <= 2; j ++{
		Clear_123()
		if j == 0 {
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Println()
		} else if j == 1 {
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Println()
		} else {
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Print("*")
			time.Sleep(500 * time.Millisecond)
			fmt.Println()
		}
	}
}