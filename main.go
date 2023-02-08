// 052-The-While-Loop.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	// get a random number between 1 to 1001
	rand.Seed(time.Now().UnixNano())

	i := 1000
	for i > 100 {
		// this is infinite loop
		i = rand.Intn(1000) + 1
		fmt.Println("i:",i)
		if i > 100 {
			fmt.Println("i:", i, "loop keeps going...")
		}
	}
}
