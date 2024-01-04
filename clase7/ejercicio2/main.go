package main

import (
	"errors"
	"fmt"
)

type salaryError struct {
}

func (e *salaryError) Error() string {
	return fmt.Sprintf("Error: salary is less than 10000")
}

func checkSalary(salary int) error {
	if salary <= 100000 {
		return &salaryError{}
	}
	return nil
}

func main() {
	var salary int
	salary = 100000
	err := checkSalary(salary)
	if errors.Is(err, &salaryError{}) {
		fmt.Println(err)
	}
}
