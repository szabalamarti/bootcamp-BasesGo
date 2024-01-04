package main

import (
	"fmt"
)

func checkSalary(salary int) error {
	if salary <= 100000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}
	return nil
}

func main() {
	var salary int
	salary = 99000
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err)
	}
}
