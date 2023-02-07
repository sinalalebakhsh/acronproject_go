package staff

import "fmt"

var OverPaidLimit = 80000
var UnderPaidLimit = 10000



type Employee struct {
	FirstName string
	LastName string
	Salary int
	FullTime bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee{
	var underpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < UnderPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	e.notVisible()
	return underpaid
}


func (e *Office) notVisible() {
	fmt.Println("It's Unexported function.")
}

/*
RESULT is:
-----
sina-lalebakhsh@sinalalebakhsh:~/Desktop/acronproject_Go$ go run main.go 
Overpaid staff: [{Mahsa Gilani 7543211 false}]
It's Unexported function.
Underpaid staff: [{Sina Lalebakhsh 400 true} {Gandom Irani 1233 false}]
sina-lalebakhsh@sinalalebakhsh:~/Desktop/acronproject_Go$ 
*/ 