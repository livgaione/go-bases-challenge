package routes

import (
	"net/http"
	"tickets/controller"

	"github.com/go-chi/chi/v5"
)

func NewRouter(ctrl controller.TicketsController) http.Handler {
	r := chi.NewRouter()

	r.Get("/tickets/{country}", ctrl.GetByCountryHandler)
	r.Get("/tickets/{time}", ctrl.GetCountByPeriod)
	r.Get("/tickets/{country}/percentage", ctrl.AverageDestination)

	return r
}
