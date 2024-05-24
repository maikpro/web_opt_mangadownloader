package routes

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/maikpro/web_opt_mangadownloader/docs"

	"github.com/maikpro/web_opt_mangadownloader/controllers"
	"github.com/maikpro/web_opt_mangadownloader/util"
	httpSwagger "github.com/swaggo/http-swagger"
)

const root = "/api"

func HandleHttp() {
	mux := http.NewServeMux()
	handler := enableCors(mux)

	// Serve Swagger UI
	mux.HandleFunc("/", redirectHandler)
	mux.HandleFunc("/api/", redirectHandler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc(fmt.Sprintf("%s/health", root), controllers.GetHealthCheck)

	mux.HandleFunc(fmt.Sprintf("%s/arcs", root), controllers.GetArcs)

	mux.HandleFunc(fmt.Sprintf("%s/chapters/id/", root), controllers.DownloadChapter)

	mux.HandleFunc(fmt.Sprintf("%s/settings", root), controllers.SettingsHandler)
	mux.HandleFunc(fmt.Sprintf("%s/settings/id/", root), controllers.UpdateSettings)

	port, err := util.GetEnvString("SERVER_PORT")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Server runs on port :%s/api/\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin, err := util.GetEnvString("ORIGIN_URL")
		if err != nil {
			log.Fatal(err)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, %s, %s", http.MethodGet, http.MethodPost, http.MethodPut))
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/", http.StatusPermanentRedirect)
}
