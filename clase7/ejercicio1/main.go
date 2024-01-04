package main

import "fmt"

type TaxError struct {
}

func (e *TaxError) Error() string {
	return fmt.Sprintf("Error: the salary entered does not reach the taxable minimum")
}

func taxSalary(salary int) (string, error) {
	if salary < 150000 {
		return "", &TaxError{}
	}
	return "Must pay tax", nil
}

func main() {
	var salary int
	salary = 100000
	msg, err := taxSalary(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
