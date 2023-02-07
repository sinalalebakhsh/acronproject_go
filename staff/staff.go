package staff

var overPaidLimit = 75000

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
		if x.Salary > overPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}