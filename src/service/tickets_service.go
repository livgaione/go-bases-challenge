package service

import (
	"tickets/domain/model"
	"tickets/repository"
)

type TicketsService interface {
	GetByCountry(destination string) ([]model.Ticket, error)
	GetCountByPeriod(time string) ([]model.Ticket, error)
	AverageDestination(destination string) ([]model.Ticket, float64, error)
}

type ticketsService struct {
	repo repository.TicketRepository
}

func NewTicketsService(r repository.TicketRepository) TicketsService {
	return &ticketsService{repo: r}
}

func (t ticketsService) GetByCountry(destination string) ([]model.Ticket, error) {
	return t.repo.GetByCountry(destination)
}

func (t ticketsService) GetCountByPeriod(time string) ([]model.Ticket, error) {
	return t.repo.GetCountByPeriod(time)
}

func (t ticketsService) AverageDestination(destination string) ([]model.Ticket, float64, error) {
	return t.repo.AverageDestination(destination)
}
