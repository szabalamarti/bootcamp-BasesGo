package tickets

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	FlightTime  time.Time
	Price       int
}

type TicketService struct {
	Tickets []Ticket
}

// NewTicketService returns a new TicketService
func NewTicketService() *TicketService {
	return &TicketService{}
}

// checkTimePeriod checks if a time is within one of four time periods
func checkTimePeriod(t time.Time, timePeriod string) bool {
	timeMap := map[string][]time.Time{
		"morning":   {time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 6, 59, 59, 0, time.UTC)},
		"noon":      {time.Date(0, 0, 0, 7, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 12, 59, 59, 0, time.UTC)},
		"afternoon": {time.Date(0, 0, 0, 13, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 19, 59, 59, 0, time.UTC)},
		"night":     {time.Date(0, 0, 0, 20, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 23, 59, 59, 0, time.UTC)},
	}

	// Create a new time.Time value for t that has the same date as the start and end times
	t = time.Date(0, 0, 0, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	timeRange := timeMap[timePeriod]
	startTime := timeRange[0]
	endTime := timeRange[1]

	// Check if t is within the time range
	return (t.Equal(startTime) || t.After(startTime)) && t.Before(endTime)
}

// LoadTickets loads tickets from a CSV file
func (ts *TicketService) LoadTickets(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		// Parse id field to int
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		// Parse price field to int
		price, err := strconv.Atoi(fields[5])
		if err != nil {
			return err
		}
		// Parse flight time field to time.Time
		flightTime, err := time.Parse("15:04", fields[4])
		if err != nil {
			return err
		}
		ticket := Ticket{
			Id:          id,
			Name:        fields[1],
			Email:       fields[2],
			Destination: fields[3],
			FlightTime:  flightTime,
			Price:       price,
		}
		// Append ticket to slice in service
		ts.Tickets = append(ts.Tickets, ticket)
	}
	return nil
}

// GetTotalTickets returns the total number of tickets for a destination
func (ts *TicketService) GetTotalTickets(destination string) (int, error) {
	var total int
	for _, ticket := range ts.Tickets {
		if ticket.Destination == destination {
			total++
		}
	}
	if total == 0 {
		return 0, errors.New("no tickets found for destination")
	}
	return total, nil
}

// GetCountByPeriod returns the total number of tickets for a time period
func (ts *TicketService) GetCountByPeriod(time string) (int, error) {
	if time != "morning" && time != "noon" && time != "afternoon" && time != "night" {
		return 0, errors.New("invalid time period, must be morning, noon, afternoon or night")
	}
	var total int
	for _, ticket := range ts.Tickets {
		if checkTimePeriod(ticket.FlightTime, time) {
			total++
		}
	}
	return total, nil
}

// AverageDestination returns the percentage of total tickets for a destination
func (ts *TicketService) AverageDestination(destination string) (float64, error) {
	total := len(ts.Tickets)
	count, err := ts.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	return float64(count) / float64(total), nil
}
