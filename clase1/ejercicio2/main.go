package main

import "fmt"

func main() {
	var temperature int64
	var humidity float64
	var pressure float64
	temperature = 16
	humidity = 0.85
	pressure = 1013.25

	fmt.Println("Temperature: ", temperature)
	fmt.Println("Humidity:", humidity)
	fmt.Println("Pressure:", pressure)
}
