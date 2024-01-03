package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Identity Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Position:", e.Position)
	fmt.Println("Name:", e.Identity.Name)
	fmt.Println("FDN:", e.Identity.DateOfBirth)
}

func main() {
	person := Person{
		ID:          1,
		Name:        "Pepito",
		DateOfBirth: "30/01/1999",
	}
	employee := Employee{
		ID:       1,
		Position: "Software Developer",
		Identity: person,
	}
	employee.PrintEmployee()
}
