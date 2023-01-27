package data

import (
	"fmt"
	"time"
	"github.com/jalaali/go-jalaali"
)

var Name string = "Sina"

func PrintNowDate()  {
	fmt.Println(time.Now().Date())
	fmt.Println(jalaali.ToJalaali(time.Now().Date()))
}