package main

import (
	"fmt"
	"time"
	"github.com/jalaali/go-jalaali"
)

func main()  {
	fmt.Println(time.Now().Date())
	fmt.Println(jalaali.ToJalaali(time.Now().Date()))
}

