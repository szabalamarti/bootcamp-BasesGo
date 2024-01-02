package main

import "fmt"

func main() {
	salary := 150000.0
	fmt.Println(getTax(salary))
}

func getTax(salary float64) (tax float64) {
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	}
	return tax
}
