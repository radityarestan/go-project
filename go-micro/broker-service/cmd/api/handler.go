package main

import (
	"log"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	log.Println("hitting the broker")
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}
