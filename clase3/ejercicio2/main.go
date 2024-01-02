package main

import "fmt"

func main() {
	average := Average(2, 3, 5, 6, 4, 1)
	fmt.Println(average)
}

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
