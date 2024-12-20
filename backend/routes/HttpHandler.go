package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	_ "github.com/maikpro/web_opt_mangadownloader/docs"

	"github.com/maikpro/web_opt_mangadownloader/controllers"
	"github.com/maikpro/web_opt_mangadownloader/util"
	httpSwagger "github.com/swaggo/http-swagger"
)

const root = "/api"

func HandleHttp() {
	router := httprouter.New()
	handler := enableCors(router)

	// Serve Swagger UI
	router.GET("/", redirectHandler)
	router.GET("/api/", redirectHandler)
	router.GET("/swagger/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	})

	router.GET(fmt.Sprintf("%s/health", root), controllers.GetHealthCheck)

	router.GET(fmt.Sprintf("%s/arcs", root), controllers.GetArcs)

	router.POST(fmt.Sprintf("%s/chapters/id/:id", root), controllers.DownloadChapter)
	router.GET(fmt.Sprintf("%s/chapters/id/:id", root), controllers.ViewChapterPage)

	router.GET(fmt.Sprintf("%s/settings", root), controllers.GetSettings)
	router.POST(fmt.Sprintf("%s/settings", root), controllers.SaveSettings)
	router.PUT(fmt.Sprintf("%s/settings/id/:id", root), controllers.UpdateSettings)

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

		if !strings.HasPrefix(r.URL.Path, "/api/chapters/id/") {
			w.Header().Set("Content-Type", "application/json")
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, %s, %s", http.MethodGet, http.MethodPost, http.MethodPut))

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func redirectHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/swagger/", http.StatusPermanentRedirect)
}
