package main

import (
	"ejercicio3/internal/salary"
	"fmt"
)

func main() {
	value := salary.GetSalary(10, "C")
	fmt.Println(value)
}
