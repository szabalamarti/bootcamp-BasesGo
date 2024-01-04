package main

import (
	"errors"
	"fmt"
)

var ErrSalaryTooLow error = errors.New("Error: salary is less than 10000")

func checkSalary(salary int) error {
	if salary <= 100000 {
		return ErrSalaryTooLow
	}
	return nil
}

func main() {
	var salary int
	salary = 99000
	err := checkSalary(salary)
	if errors.Is(err, ErrSalaryTooLow) {
		fmt.Println(err)
	}
}
