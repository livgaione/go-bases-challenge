package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tickets/domain/loader"
	"tickets/domain/model"
)

type TicketRepository interface {
	GetByCountry(destination string) ([]model.Ticket, error)
	GetCountByPeriod(time string) ([]model.Ticket, error)
	AverageDestination(destination string) ([]model.Ticket, float64, error)
}

type ticketRepository struct {
	tickets []model.Ticket
}

func NewRepository() TicketRepository {
	tickets := loader.LoadTickets()
	return &ticketRepository{tickets: tickets}
}

func (t ticketRepository) GetByCountry(destination string) ([]model.Ticket, error) {
	if destination == "" {
		return nil, errors.New("destination cannot be empty")
	}

	var result []model.Ticket
	for _, ticket := range t.tickets {
		if ticket.Country == destination {
			result = append(result, ticket)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no tickets found for destination: %s", destination)
	}

	return result, nil
}

func (t ticketRepository) GetCountByPeriod(time string) ([]model.Ticket, error) {

	validPeriods := map[string]bool{
		"Dawn":      true,
		"Morning":   true,
		"Afternoon": true,
		"Evening":   true,
	}

	if !validPeriods[time] {
		return nil, fmt.Errorf("invalid time period: %s", time)
	}

	var result []model.Ticket
	for _, tickets := range t.tickets {

		h := strings.Split(tickets.Hour, ":")

		if len(h) < 1 {
			continue
		}

		hour, err := strconv.Atoi(h[0])
		if err != nil {
			continue
		}

		switch time {
		case "Dawn":
			if hour >= 0 && hour <= 6 {
				result = append(result, tickets)
			}
		case "Morning":
			if hour >= 7 && hour <= 12 {
				result = append(result, tickets)
			}
		case "Afternoon":
			if hour >= 13 && hour <= 18 {
				result = append(result, tickets)
			}
		case "Evening":
			if hour >= 19 && hour <= 23 {
				result = append(result, tickets)
			}
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no tickets found for time period: %s", time)
	}

	return result, nil
}

func (t ticketRepository) AverageDestination(destination string) ([]model.Ticket, float64, error) {
	if destination == "" {
		return nil, 0, errors.New("destination cannot be empty")
	}

	var result []model.Ticket
	total := len(t.tickets)
	if total == 0 {
		return nil, 0, errors.New("no tickets available to calculate average")
	}

	count := 0
	for _, ticket := range t.tickets {
		if ticket.Country == destination {
			result = append(result, ticket)
			count++
		}
	}

	if count == 0 {
		return nil, 0, fmt.Errorf("no tickets found for destination: %s", destination)
	}

	percentage := (float64(count) / float64(total)) * 100
	return result, percentage, nil
}
