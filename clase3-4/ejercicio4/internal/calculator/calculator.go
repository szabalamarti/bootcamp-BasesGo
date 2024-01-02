package calculator

import "errors"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Operation(operation_type string) (func(values ...int) float64, error) {
	switch operation_type {
	case minimum:
		function := Minimum
		return function, nil
	case maximum:
		function := Maximum
		return function, nil
	case average:
		function := Average
		return function, nil
	default:
		return nil, errors.New("unsopported operand")
	}
}
