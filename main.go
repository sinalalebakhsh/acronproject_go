//043-Compound-Booleans.go
package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	// boolean expression
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > < >= <=
	fmt.Printf("%d > %d: %t \n", apples, oranges, apples > oranges)
	fmt.Printf("%d < %d: %t \n", apples, oranges, apples < oranges)
	fmt.Printf("%d >= %d: %t \n", apples, oranges, apples >= oranges)
	fmt.Printf("%d <= %d: %t \n", apples, oranges, apples <= oranges)
}
