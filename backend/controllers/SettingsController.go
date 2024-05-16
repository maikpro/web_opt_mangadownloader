package controllers

import (
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
}
