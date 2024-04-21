package employee

// EmployeeInterface is an interface for employee
type EmployeeInterface interface {
	SetName(name string)
	SetSalary(salary int)
	GetName() string
	GetSalary() int
	/*
		TODO: Add a new method called GetBonus() that returns a float64
	*/
	GetBonus() float64
}

// Employee is a struct for employee
type Employee struct {
	Name   string
	Salary int
	Bonus  float64
}

// SetName sets the name of the employee
func (e *Employee) SetName(name string) {
	// TODO: Set the name of the employee
	e.Name = name
}

// SetSalary sets the salary of the employee
func (e *Employee) SetSalary(salary int) {
	// TODO: Set the salary of the employee
	e.Salary = salary
}

// GetName gets the name of the employee
func (e *Employee) GetName() string {
	// TODO: Get the name of the employee
	return e.Name
}

// GetSalary gets the salary of the employee
func (e *Employee) GetSalary() int {
	// TODO: Get the salary of the employee
	return e.Salary
}

func (e *Employee) GetBonus() float64 {
	return e.Bonus
}