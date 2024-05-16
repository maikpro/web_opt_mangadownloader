package controllers

import (
	"io"
	"log"
	"net/http"
)

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("That's not a POST Request!")
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Print the received data
	log.Println("Received data:", string(body))

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}
