package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	service := tickets.NewTicketService()
	err := service.LoadTickets("tickets.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	destinationCountry := "Argentina"
	period := "morning"

	// Exercise 1
	totalDestination, err := service.GetTotalTickets(destinationCountry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total tickets for %s: %d\n", destinationCountry, totalDestination)

	// Exercise 2
	periodCount, err := service.GetCountByPeriod(period)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total tickets for %s: %d\n", period, periodCount)

	// Exercise 3
	averageDestination, err := service.AverageDestination(destinationCountry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Percentage of total tickets for %s: %f\n", destinationCountry, averageDestination)
}
