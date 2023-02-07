//047-Exported-VS-Unexported.go
package main

import (
	"acron/staff"
	"fmt"
)

func main() {
	myStaff := staff.Office {
		AllStaff: employees,
	}
	// fmt.Println(myStaff.All())
	staff.OverPaidLimit = 8000000
	fmt.Println(myStaff.Overpaid())
}

var employees = []staff.Employee {
	{FirstName: "Sina", LastName: "Lalebakhsh", Salary: 400, FullTime: true,},
	{FirstName: "Gandom", LastName: "Irani", Salary: 1233, FullTime: false,},
	{FirstName: "Kimia", LastName: "Tehrani", Salary: 12345, FullTime: false,},
	{FirstName: "Mahsa", LastName: "Gilani", Salary: 7543211, FullTime: false,},
}