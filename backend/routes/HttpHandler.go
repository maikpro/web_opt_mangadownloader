package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maikpro/web_opt_mangadownloader/controllers"
)

const root = "/api"

func HandleHttp() {
	mux := http.NewServeMux()
	handler := enableCors(mux)

	mux.HandleFunc(fmt.Sprintf("%s/arcs", root), controllers.GetArcs)
	mux.HandleFunc(fmt.Sprintf("%s/chapters/id/", root), controllers.DownloadChapter)
	mux.HandleFunc(fmt.Sprintf("%s/settings", root), controllers.SaveSettings)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}
