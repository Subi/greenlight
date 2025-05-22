package main

import (
	"encoding/json"
	"net/http"
)

var ErrorProcessing = "Server encountered an error trying to process the request"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	v, err := json.Marshal(data)
	if err != nil {
		app.logger.Error("Error marshalling data", err)
		http.Error(w, ErrorProcessing, http.StatusInternalServerError)
		return
	}

	// Add a new line to JSON for better readability in terminals
	v = append(v, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}
