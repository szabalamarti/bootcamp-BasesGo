package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	// Age of Benjamin
	nombre := "Benjamin"
	fmt.Println("Edad de Benjamin:", employees[nombre])

	// Employees older than 21
	total := 0
	for _, age := range employees {
		if age > 21 {
			total += 1
		}
	}
	fmt.Println("Total de empleados mayores de 21:", total)

	// Add Federico, age 25, to the map
	employees["Federico"] = 25

	// Delete Pedro
	delete(employees, "Pedro")

	// Print final map
	fmt.Println("\nMapeo final:")
	for employee, age := range employees {
		fmt.Println(employee, age)
	}
}
