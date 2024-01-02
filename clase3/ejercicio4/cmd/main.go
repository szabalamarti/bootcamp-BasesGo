package main

import (
	"ejercicio4/internal/calculator"
	"fmt"
)

func main() {
	minFunc, minErr := calculator.Operation("minimum")
	averageFunc, avgErr := calculator.Operation("average")
	maxFunc, maxErr := calculator.Operation("maximum")
	if minErr != nil {
		panic(minErr)
	}
	if avgErr != nil {
		panic(avgErr)
	}
	if maxErr != nil {
		panic(maxErr)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println("Minimun value is: ", minValue)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Average value is: ", averageValue)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Maximum value is: ", maxValue)
}
