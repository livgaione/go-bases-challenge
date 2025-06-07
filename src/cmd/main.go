package main

import (
	"log"
	"net/http"
	"tickets/controller"
	"tickets/repository"
	"tickets/routes"
	"tickets/service"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewTicketsService(repo)
	ctrl := controller.TicketsController{Service: svc}

	r := routes.NewRouter(ctrl)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
