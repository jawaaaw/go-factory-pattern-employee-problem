package employee

// Director is a struct for Director
// It embeds Employee
// This means that Director has all the fields and methods of Employee
type Director struct {
	// TODO: Add the Employee struct
	Employee
}

// NewDirector creates a new director
// It returns a pointer to the director
// Creational method
func NewDirector() *Director {
	// TODO: Create a new director
	// Set the name to "Director"
	// Set the salary to 5000
	director := & Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  5000.0 * 0.3,
		},
	}
	return director
}