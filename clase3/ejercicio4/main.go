package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, minErr := operation(minimum)
	averageFunc, avgErr := operation(average)
	maxFunc, maxErr := operation(maximum)
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

func operation(operation_type string) (func(values ...int) float64, error) {
	switch operation_type {
	case minimum:
		function := func(values ...int) float64 {
			min := values[0]
			for i := 1; i < len(values); i++ {
				if min > values[i] {
					min = values[i]
				}
			}
			return float64(min)
		}
		return function, nil
	case maximum:
		function := func(values ...int) float64 {
			max := values[0]
			for i := 1; i < len(values); i++ {
				if max < values[i] {
					max = values[i]
				}
			}
			return float64(max)
		}
		return function, nil
	case average:
		function := func(values ...int) float64 {
			var total int
			lenght := len(values)
			for _, grade := range values {
				total += grade
			}
			if lenght == 0 {
				return 0
			}
			return float64(total) / float64(lenght)
		}
		return function, nil
	default:
		return nil, errors.New("Unsopported operand")
	}
}
