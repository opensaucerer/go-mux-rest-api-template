package controller

import (
	"net/http"

	"github.com/opensaucerer/gotemplate/helper"
	"github.com/opensaucerer/gotemplate/typing"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := typing.Health{
		Version:     "0.0.1",
		Status:      true,
		Description: "Go Template is up and running. This repository is to serve as a template for future Go projects.",
	}
	helper.SendJSONResponse(w, true, http.StatusOK, "Health Check", data)
}
