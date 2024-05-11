package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maikpro/web_opt_mangadownloader/services"
)

func GetArcs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("That's not a GET Request!")
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	arcs, err := services.GetArcList()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(arcs)
}
