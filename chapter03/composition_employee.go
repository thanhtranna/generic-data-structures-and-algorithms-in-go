package main

import "fmt"

type employee struct {
	lastName  string
	firstName string
	role      string
	salary    float64
}

type Employee interface {
	SetLastName(lName string)
	SetFirstName(fName string)
	SetRole(r string)
	SetSalary(s float64)
	GetSalary() float64
	String() string
}

type partTimeEmployee struct {
	employee
	hourlyWage float64
}

type PartTimeEmployee interface {
	Employee
	SetHourlyWage(hourly float64)
	GetHourlyWage() float64
}

// Methods
func (person *employee) SetSalary(yearly float64) {
	person.salary = yearly
}

func (person employee) GetSalary() float64 {
	return person.salary
}

func (person *employee) SetFirstName(firstN string) {
	person.firstName = firstN
}

func (person employee) GetFirstName() string {
	return person.firstName
}

func (person *employee) SetLastName(lastN string) {
	person.lastName = lastN
}

func (person *employee) SetRole(r string) {
	person.role = r
}

func (person employee) GetRole() string {
	return person.role
}

func (person employee) String() string {
	result := "Name: " + person.firstName + " " + person.lastName + "\n"
	result += "Role: " + person.role + "\n"
	result += "Annual salary: $" + fmt.Sprintf("%0.2f", person.
		salary) + "\n"
	return result
}

func (person partTimeEmployee) String() string {
	result := "Name: " + person.firstName + " " + person.lastName + "\n"
	result += "Role: " + person.role + "\n"
	result += "HourlyWage: $" + fmt.Sprintf("%0.2f", person.
		hourlyWage) + "\n"
	return result
}

func (person *partTimeEmployee) SetHourlyWage(amt float64) {
	person.hourlyWage = amt
}

func (person partTimeEmployee) GetHourlyWage() float64 {
	return person.hourlyWage
}

func main() {
	person := new(employee) // Returns the address of an employee
	person.SetFirstName("Helen")
	person.SetLastName("Rose")
	person.SetRole("Technical Lead")
	person.SetSalary(125_644.0)
	fmt.Println(person.String())

	hourlyWorker := new(partTimeEmployee) // Returns address
	hourlyWorker.SetFirstName("Mark")
	hourlyWorker.SetLastName("Smith")
	hourlyWorker.SetRole("Software Developer")
	hourlyWorker.SetHourlyWage(85.00)
	fmt.Println(hourlyWorker.String())
}
