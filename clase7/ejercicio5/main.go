package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hoursWorked int, hourlySalary float64) (float64, error) {
	if hoursWorked < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	var salary float64 = float64(hoursWorked) * hourlySalary
	if salary >= 150000.0 {
		salary *= 0.9
	}

	return salary, nil
}

func main() {
	var hoursWorked int
	var hourlySalary float64
	hoursWorked = 60
	hourlySalary = 5.0
	fmt.Println("Worked", hoursWorked, "hours at", hourlySalary, "$ the hour")
	salary, err := calculateSalary(hoursWorked, hourlySalary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Salary is: ", salary)

}
