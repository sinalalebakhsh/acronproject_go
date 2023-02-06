//043-Compound-Booleans.go
package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack",
		Age:      40,
		Salary:   2000000,
		FullTime: false,
	}

	sina := Employee{
		Name:     "Sina",
		Age:      29,
		Salary:   20000000000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, sina)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older.")
		} else {
			fmt.Println(x.Name, "is under 30.")
		}

		if x.Age < 35 && x.Salary > 50000 {
			fmt.Println(x.Name, "is under 35 and makes more than 50,000$")
		} else {
			fmt.Println(x.Name, "is over 35 Or makes less than 50,000$")
		}

		if (x.Age > 35 || x.Salary < 5000) && x.FullTime == false {
			fmt.Println("\n" + x.Name, "is more than 35 age OR\nmakes money less than 5,000$ and doesn't Full Time job!!!\nSo It's better to get a Full Time job.\n")
		}

	}

}
