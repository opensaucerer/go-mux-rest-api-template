package version

import (
	mux "github.com/gorilla/mux"
	routes "github.com/opensaucerer/gotemplate/route"
)

func VersionRoutes(r *mux.Router) {
	router := r.PathPrefix("/v1").Subrouter()

	routes.RegisterBaseRoutes(router)
}
