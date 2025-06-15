package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/maikpro/web_opt_mangadownloader/bootstrap"
	"github.com/maikpro/web_opt_mangadownloader/services"

	"github.com/maikpro/web_opt_mangadownloader/api/controllers"
)

const root = "/api"

func HandleHttp(env *bootstrap.Env) {
	router := httprouter.New()
	handler := enableCors(env, router)

	// HealthController
	router.GET(fmt.Sprintf("%s/health", root), controllers.GetHealthCheck)

	// OPTClient dependency
	var optClient services.IOPTClient = &services.OPTClient{}

	// ArcController
	arcController := controllers.ArcController{OptClient: optClient}
	router.GET(fmt.Sprintf("%s/arcs", root), arcController.GetArcs)

	// ChapterController
	chapterContoller := controllers.ChapterController{OptClient: optClient}
	router.GET(fmt.Sprintf("%s/chapters/id/:id", root), chapterContoller.GetChapter)
	router.POST(fmt.Sprintf("%s/chapters/id/:id", root), chapterContoller.DownloadChapter)

	// SettingsController
	router.GET(fmt.Sprintf("%s/settings", root), controllers.GetSettings)
	router.POST(fmt.Sprintf("%s/settings", root), controllers.SaveSettings)
	router.PUT(fmt.Sprintf("%s/settings/id/:id", root), controllers.UpdateSettings)

	log.Printf("Server runs on port :%s/api/\n", env.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), handler)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func enableCors(env *bootstrap.Env, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/chapters/id/") {
			w.Header().Set("Content-Type", "application/json")
		}

		w.Header().Set("Access-Control-Allow-Origin", env.OriginUrl)
		w.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, %s, %s", http.MethodGet, http.MethodPost, http.MethodPut))

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
