// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package main

import (
	"fmt"
	"acron/con_cap"
)


func main() {
	for key, value := range con_cap.CountriesCapital {
		fmt.Println(key, "	:	", value)
	}

}



