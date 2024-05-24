package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maikpro/web_opt_mangadownloader/services"
)

// Responds to a HTTP GET Request with all Arcs + Chapters fetched from OnePiece-Tube.com
// @Summary fetches Arcs from OPT
// @Description delivers Arcs from OPT
// @Tags arcs
// @Accept json
// @Produce json
// @Success 200 {object} []models.Arc
// @Router /api/arcs [get]
func GetArcs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("That's not a GET Request!")
		http.NotFound(w, r)
		return
	}

	arcs, err := services.GetArcList()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(arcs)
}
