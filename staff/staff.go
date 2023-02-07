package staff

type Employee struct {
	FirstName string
	LastName string
	Salary int
	FullTime bool
}

type Office struct {
	AllStaff []Employee
}