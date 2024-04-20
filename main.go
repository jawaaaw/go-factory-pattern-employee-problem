package main

import (
	"fmt"
	"go-factorypattern-company-case/employee"
)
type EmployeeInterface interface {
    GetName() string
    GetSalary() int
	CalculateBonus() int
}

type Manager struct {
    Name   string
    Salary int
}

func (m *Manager) GetName() string {
    return m.Name
}

func (m *Manager) GetSalary() int {
    return m.Salary
}

func (m Manager) CalculateBonus() int {
	return int(0.2 * float64(m.Salary)) // 20% bonus for managers
}

type Developer struct {
    Name   string
    Salary int
}

func (d *Developer) GetName() string {
    return d.Name
}

func (d *Developer) GetSalary() int {
    return d.Salary
}

func (d *Developer) CalculateBonus() int {
	// Calculate the bonus for a developer (e.g., 15% of salary)
	return int(0.15 * float64(d.Salary))
}

func GetEmployeeFactory(empType string) (EmployeeInterface, error) {
    switch empType {
    case "manager":
        return &Manager{Name: "Budi", Salary: 5000}, nil
    case "developer":
        return &Developer{Name: "Joko", Salary: 4000}, nil
    default:
        return nil, fmt.Errorf("unknown employee type: %s", empType)
    }
}
// use main function to demonstrate the factory pattern
func main() {

	budi, err := employee.GetEmployeeFactory("manager")
	if err != nil {
		panic(err)
	}

	printDetails(budi)

}

func printDetails(e employee.EmployeeInterface) {
	fmt.Printf("Employee: %s", e.GetName())
	fmt.Println()
	fmt.Printf("Salary: %d", e.GetSalary())
	fmt.Println()
}
