package repository

import (
	"testing"

	"tickets/domain/model"

	"github.com/stretchr/testify/require"
)

func TestGetByCountry(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, _ := repo.GetByCountry("Brazil")
	require.Len(t, result, 2)
	require.Equal(t, "Brazil", result[0].Country)
	require.Equal(t, "Brazil", result[1].Country)

}

func TestGetByCountryWithNoDestination(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, err := repo.GetByCountry("")

	require.Error(t, err)
	require.Nil(t, result)

}
func TestGetByCountryWithNoFoundDestination(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, err := repo.GetByCountry("China")

	require.Error(t, err)
	require.Nil(t, result)

}

func TestGetCountByPeriod(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, _ := repo.GetCountByPeriod("Morning")
	require.Len(t, result, 2)
	require.Equal(t, "Brazil", result[0].Country)
	require.Equal(t, "Argentina", result[1].Country)

}

func TestGetCountByPeriodWithInvalidPeriod(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, err := repo.GetCountByPeriod("Invalid")
	require.Error(t, err)
	require.Nil(t, result)
}

func TestGetCountByPeriodWithNoFoundTickers(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, err := repo.GetCountByPeriod("China")
	require.Error(t, err)
	require.Nil(t, result)
}

func TestAverageDestination(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
		{Country: "China", Hour: "18:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, count, _ := repo.AverageDestination("Argentina")
	require.Len(t, result, 1)
	require.Equal(t, 25.0, count)

}

func TestAverageDestinationWithNoDestionation(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
		{Country: "China", Hour: "18:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, _, err := repo.AverageDestination("")
	require.Error(t, err)
	require.Nil(t, result)

}

func TestAverageDestinationWithNoFoundDestionation(t *testing.T) {

	mockTickets := []model.Ticket{
		{Country: "Brazil", Hour: "10:00"},
		{Country: "Argentina", Hour: "12:00"},
		{Country: "Brazil", Hour: "15:00"},
	}

	repo := ticketRepository{tickets: mockTickets}

	result, _, err := repo.AverageDestination("Colombia")
	require.Error(t, err)
	require.Nil(t, result)

}
