// 054-Nested-Loops-and-the-Debugger.go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print("i:", i)
		for j := 1; j<= 3; j++ {
			fmt.Print("    j: ", j)
		}
		fmt.Println()
	}
}


/*
I learn , Working with Debugger in this episode. It's very useful. 
*/ 