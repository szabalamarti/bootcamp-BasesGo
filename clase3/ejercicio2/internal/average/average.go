package average

func Average(grades ...int) (average float64) {
	var total int
	lenght := len(grades)
	for _, grade := range grades {
		total += grade
	}
	if lenght == 0 {
		return 0
	}
	return float64(total) / float64(lenght)
}
