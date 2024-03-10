package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	// ignore the error
	_ = app.writeJSON(w, http.StatusOK, payload)
	// The payload struct is then encoded into JSON format using json.MarshalIndent() function,
	// which returns the JSON representation of the payload. The out variable holds the JSON-encoded data.
}
