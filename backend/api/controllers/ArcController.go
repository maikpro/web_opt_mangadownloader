package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maikpro/web_opt_mangadownloader/services"
)

type ArcController struct {
	OptClient services.IOPTClient
}

func (arcController *ArcController) GetArcs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		log.Println("That's not a GET Request!")
		http.NotFound(w, r)
		return
	}

	arcs, err := arcController.OptClient.GetArcList()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(arcs)
}
