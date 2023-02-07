//047-Exported-VS-Unexported.go
package main

import "acron/staff"

func main() {

}

var employees = []staff.Employee {
	{FirstName: "Sina", LastName: "Lalebakhsh", Salary: 900000000, FullTime: true,},
	{FirstName: "Gandom", LastName: "Irani", Salary: 123000000, FullTime: false,},
	{FirstName: "Kimia", LastName: "Tehrani", Salary: 4302000000, FullTime: false,},
	{FirstName: "Mahsa", LastName: "Gilani", Salary: 999999999999, FullTime: false,},
}