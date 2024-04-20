package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	// TODO: Add the Employee struct
	Employee
}

// NewIntern creates a new intern
func NewIntern() *Intern {
	// TODO: Create a new intern
	// Set the name to "Intern"
	// Set the salary to 100
	intern := &Intern{
		Employee: Employee{
			Name:   "Intern",
			Salary: 100,
		},
	}
	return intern
}

func (i *Intern) GetBonus() float64 {
	return float64(i.Salary) * 0.0
}