// 053-The-Infinite-Loop.go
package main

import (
	"acron/mylogger"
	"bufio"
	"fmt"
	"os"
	"time"
)


func main() {
	// read input from the user 5 times and write it to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)			

	// Get log from user,  concurrent program
	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something:")
	for i:= 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
