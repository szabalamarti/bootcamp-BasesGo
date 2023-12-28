package main

import "fmt"

func main() {
	// Variables
	var age int = 23
	var employed bool = true
	var seniority float64 = 2
	var salary = 100001

	if age <= 22 || employed == false || seniority <= 1.0 {
		fmt.Println("Prestamo denegado")
	} else {
		fmt.Println("Prestamo otorgado")
		if salary > 100000 {
			fmt.Println("No se cobra interes")
		}
	}
}
