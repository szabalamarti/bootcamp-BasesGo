package main

import "fmt"

const (
	categoryA = "A"
	categoryB = "B"
	categoryC = "C"
)

func main() {
	fmt.Println(getSalary(10, "C"))
}

func getSalary(minutes int, category string) (salary int) {
	var monthly_salary float64
	switch category {
	case categoryA:
		monthly_salary = 1000.0 * float64(minutes) / 60.0
	case categoryB:
		monthly_salary = 1500 * float64(minutes) / 60
		monthly_salary *= 1.2
	case categoryC:
		monthly_salary = 3000 * float64(minutes) / 60
		monthly_salary *= 1.5
	}
	return int(monthly_salary)
}
