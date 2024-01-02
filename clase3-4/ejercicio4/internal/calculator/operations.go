package calculator

func Minimum(values ...int) float64 {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if min > values[i] {
			min = values[i]
		}
	}
	return float64(min)
}

func Maximum(values ...int) float64 {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if max < values[i] {
			max = values[i]
		}
	}
	return float64(max)
}

func Average(values ...int) float64 {
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
