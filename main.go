// 052-The-While-Loop.go
package main

import (
	"math/rand"
	"time"
)


func main() {
	// get a random number between 1 to 1001
	rand.Seed(time.Now().UnixNano())
	i := 1000
	for i > 100 {
		// this is infinite loop
	}
}
