package employee

import "errors"

// GetEmployeeFactory is a factory for employee
// It returns an employee based on the input
// If the employee is not found, it returns an error
func GetEmployeeFactory(empl string) (EmployeeInterface, error) {

	if empl == "manager" {
		// TODO: Return a new manager
		return NewManager(), nil
	}

	if empl == "staff" {
		// TODO: Return a new staff
		return NewStaff(), nil
	}

	if empl == "intern" {
		// TODO: Return a new intern
		return NewIntern(), nil
	}

	// TODO: Create director object
	if empl == "director" {
		return NewDirector(), nil
	}

	return nil, errors.New("Employee not found")
}