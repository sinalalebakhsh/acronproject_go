package main

import "fmt"

func main() {
	run_fib()
	convertToThird()
	fmt.Println("")

}


// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func run_fib() {
	f := fib()
	for i := 0; i<= 10-1; i++{
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Print(f(), " ")
		fmt.Println()
	}
}





func convertToThird() {
	listNumbers := []int{}
	f := fib()
	for i := 0; i<= 50-1; i++{
		listNumbers = append(listNumbers, f())
	}

	counter := 0
	for _, number := range listNumbers{
		if counter < 5 {
			defer fmt.Print(number, " ")
			counter = counter + 1
		} else {
			defer fmt.Println(" ")
			defer fmt.Print(number, " ")
			counter = 1
		}
	}
}


/* Output

1 1 2 3 5 
8 13 21 34 55 
89 144 233 377 610 
987 1597 2584 4181 6765 
10946 17711 28657 46368 75025 
121393 196418 317811 514229 832040 
1346269 2178309 3524578 5702887 9227465 
14930352 24157817 39088169 63245986 102334155 
165580141 267914296 433494437 701408733 1134903170 
1836311903 2971215073 4807526976 7778742049 12586269025 
12586269025 7778742049 4807526976 2971215073 1836311903  
1134903170 701408733 433494437 267914296 165580141  
102334155 63245986 39088169 24157817 14930352  
9227465 5702887 3524578 2178309 1346269  
832040 514229 317811 196418 121393  
75025 46368 28657 17711 10946  
6765 4181 2584 1597 987  
610 377 233 144 89  
55 34 21 13 8  
5 3 2 1 1 


*/


