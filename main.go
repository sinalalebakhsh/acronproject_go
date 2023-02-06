//039 Channels

package main

import "fmt"

// reference types (pointers, silence, maps, functions, channels)

// interface type

// RESULT => First concurrently ------
// Go routine

func main() {
	go doSomeThing("Hello, World")

	fmt.Println("This is another message")
	for  {
		// do nothing
	}
}

func doSomeThing(s string)  {
	until := 0
	for  {
		fmt.Println("s is", s)
		until = until + 1
		if until == 6 {
			break
		}
	}
}

