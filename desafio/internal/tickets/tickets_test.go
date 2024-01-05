package tickets_test

import (
	"testing"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	// Arrange
	service := &tickets.TicketService{
		Tickets: []tickets.Ticket{
			{Id: 1, Destination: "New York"},
			{Id: 2, Destination: "New York"},
			{Id: 3, Destination: "Los Angeles"},
		},
	}

	// Act and Assert
	t.Run("destination with multiple tickets", func(t *testing.T) {
		total, err := service.GetTotalTickets("New York")
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})

	t.Run("destination with one ticket", func(t *testing.T) {
		total, err := service.GetTotalTickets("Los Angeles")
		require.NoError(t, err)
		require.Equal(t, 1, total)
	})

	t.Run("destination with no tickets", func(t *testing.T) {
		total, err := service.GetTotalTickets("Chicago")
		require.Error(t, err)
		require.Equal(t, 0, total)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	// Arrange
	service := &tickets.TicketService{
		Tickets: []tickets.Ticket{
			{Id: 1, FlightTime: time.Date(2021, 1, 1, 10, 0, 0, 0, time.UTC)},
			{Id: 2, FlightTime: time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC)},
			{Id: 3, FlightTime: time.Date(2021, 1, 1, 14, 0, 0, 0, time.UTC)},
			{Id: 4, FlightTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{Id: 5, FlightTime: time.Date(2021, 1, 1, 21, 0, 0, 0, time.UTC)},
		},
	}

	// Act and Assert
	t.Run("morning", func(t *testing.T) {
		total, err := service.GetCountByPeriod("morning")
		require.NoError(t, err)
		require.Equal(t, 0, total)
	})
	t.Run("noon", func(t *testing.T) {
		total, err := service.GetCountByPeriod("noon")
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})
	t.Run("afternoon", func(t *testing.T) {
		total, err := service.GetCountByPeriod("afternoon")
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})
	t.Run("night", func(t *testing.T) {
		total, err := service.GetCountByPeriod("night")
		require.NoError(t, err)
		require.Equal(t, 1, total)
	})
	t.Run("invalid time period", func(t *testing.T) {
		total, err := service.GetCountByPeriod("brunch")
		require.Error(t, err)
		require.Equal(t, 0, total)
	})
}

func TestAverageDestination(t *testing.T) {
	// Arrange
	service := &tickets.TicketService{
		Tickets: []tickets.Ticket{
			{Id: 1, Destination: "New York"},
			{Id: 2, Destination: "New York"},
			{Id: 3, Destination: "Los Angeles"},
			{Id: 4, Destination: "Florida"},
		},
	}

	// Act and Assert
	t.Run("destination with multiple tickets", func(t *testing.T) {
		average, err := service.AverageDestination("New York")
		require.NoError(t, err)
		require.Equal(t, 0.5, average)
	})

	t.Run("destination with one ticket", func(t *testing.T) {
		average, err := service.AverageDestination("Los Angeles")
		require.NoError(t, err)
		require.Equal(t, 0.25, average)
	})

	t.Run("destination with no tickets", func(t *testing.T) {
		average, err := service.AverageDestination("Chicago")
		require.Error(t, err)
		require.Equal(t, 0.0, average)
	})
}
