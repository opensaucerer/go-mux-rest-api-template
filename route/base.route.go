package routes

import (
	controller "github.com/opensaucerer/gotemplate/controller"

	mux "github.com/gorilla/mux"
)

func RegisterBaseRoutes(r *mux.Router) {

	router := r.PathPrefix("/health").Subrouter()

	router.HandleFunc("", controller.HealthCheck).Methods("GET")
}
