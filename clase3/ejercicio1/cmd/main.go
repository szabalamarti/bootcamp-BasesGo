package main

import (
	"ejercicio1/internal/taxes"
	"fmt"
)

func main() {
	salary := 150000.0
	fmt.Println(taxes.GetTax(salary))
}
