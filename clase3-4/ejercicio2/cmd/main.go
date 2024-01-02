package main

import (
	"ejercicio2/internal/average"
	"fmt"
)

func main() {
	average := average.Average(2, 3, 5, 6, 4, 1)
	fmt.Println(average)
}
