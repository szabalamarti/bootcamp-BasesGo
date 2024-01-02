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
