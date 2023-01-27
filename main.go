package main

import (
	"fmt"
	"time"
	"github.com/sinalalebakhsh/acronproject_go/data"
	"github.com/jalaali/go-jalaali"
)

func main()  {
	fmt.Println(time.Now().Date())
	fmt.Println(jalaali.ToJalaali(time.Now().Date()))
	fmt.Println(data.Name)
}

