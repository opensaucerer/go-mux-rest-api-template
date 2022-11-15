package helper

import (
	"encoding/json"
	"net/http"

	"github.com/opensaucerer/gotemplate/typing"
)

func SendJSONResponse(w http.ResponseWriter, status bool, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(typing.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
