package controller

import (
	"encoding/json"
	"net/http"
	"tickets/service"

	"github.com/go-chi/chi/v5"
)

type TicketsController struct {
	Service service.TicketsService
}

func (tc TicketsController) GetByCountryHandler(w http.ResponseWriter, r *http.Request) {
	destination := chi.URLParam(r, "country")

	tickets, err := tc.Service.GetByCountry(destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"country": destination,
		"count":   len(tickets),
		"data":    tickets,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (tc TicketsController) GetCountByPeriod(w http.ResponseWriter, r *http.Request) {
	time := chi.URLParam(r, "time")

	tickets, err := tc.Service.GetCountByPeriod(time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"time":  time,
		"count": len(tickets),
		"data":  tickets,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (tc TicketsController) AverageDestination(w http.ResponseWriter, r *http.Request) {
	destination := chi.URLParam(r, "country")

	tickets, percentage, err := tc.Service.AverageDestination(destination)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"country":    destination,
		"percentage": percentage,
		"data":       tickets,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
