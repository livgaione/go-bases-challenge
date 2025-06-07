package service

import (
	"testing"
	"tickets/domain/model"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	GetByCountryFn       func(destination string) ([]model.Ticket, error)
	GetCountByPeriodFn   func(time string) ([]model.Ticket, error)
	AverageDestinationFn func(destination string) ([]model.Ticket, float64, error)
}

func (m *mockRepository) GetByCountry(destination string) ([]model.Ticket, error) {
	return m.GetByCountryFn(destination)
}

func (m *mockRepository) GetCountByPeriod(time string) ([]model.Ticket, error) {
	return m.GetCountByPeriodFn(time)
}

func (m *mockRepository) AverageDestination(destination string) ([]model.Ticket, float64, error) {
	return m.AverageDestinationFn(destination)
}

func TestGetByCountry(t *testing.T) {
	mock := &mockRepository{
		GetByCountryFn: func(destination string) ([]model.Ticket, error) {
			return []model.Ticket{
				{ID: "1", Country: "Brazil"},
				{ID: "2", Country: "Brazil"},
			}, nil
		},
	}

	service := NewTicketsService(mock)

	tickets, err := service.GetByCountry("Brazil")

	assert.NoError(t, err)
	assert.Len(t, tickets, 2)
	assert.Equal(t, "Brazil", tickets[0].Country)
}

func TestAverageDestination(t *testing.T) {
	mock := &mockRepository{
		AverageDestinationFn: func(destination string) ([]model.Ticket, float64, error) {
			return []model.Ticket{
				{ID: "1", Country: "Brazil"},
			}, 20.0, nil
		},
	}

	service := NewTicketsService(mock)

	tickets, percent, err := service.AverageDestination("Brazil")

	assert.NoError(t, err)
	assert.Equal(t, 20.0, percent)
	assert.Len(t, tickets, 1)
}

func TestGetCountByPeriod(t *testing.T) {
	mock := &mockRepository{
		GetCountByPeriodFn: func(time string) ([]model.Ticket, error) {
			return []model.Ticket{
				{Country: "Brazil", Hour: "10:00"},
			}, nil
		},
	}

	service := NewTicketsService(mock)

	tickets, err := service.GetCountByPeriod("Morning")

	assert.NoError(t, err)
	assert.Len(t, tickets, 1)
}
