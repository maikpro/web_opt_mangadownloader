package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/maikpro/web_opt_mangadownloader/services"
)

type SavedDirectory struct {
	ChapterName   string `json:"chapterName"`
	ChapterNumber uint   `json:"chapterNumber"`
	Path          string `json:"path"`
}

// Responds to a HTTP POST Request with a pathvariable to download specific chapter with chapterNumber
// @Summary downloads chapter from OPT
// @Description downloads chapter from OPT
// @Tags chapter
// @Accept json
// @Produce json
// @Param selectedChapter path string true "Selected Chapter"
// @Param local query string false "should download local"
// @Param telegram query string false "should send to telegram chat"
// @Success 200 {object} SavedDirectory
// @Router /api/chapters/id/{selectedChapter} [post]
func DownloadChapter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("That's not a POST Request!")
		http.NotFound(w, r)
		return
	}

	pathVariable := strings.Split(r.URL.Path, "/")
	chapterNumber := pathVariable[len(pathVariable)-1]

	if len(chapterNumber) == 0 {
		http.Error(w, "Path variable 'id' is required", http.StatusBadRequest)
		return
	}

	chapterNumberInt, err := strconv.Atoi(string(chapterNumber))
	if err != nil {
		http.Error(w, "Invalid ID format: must be an integer", http.StatusBadRequest)
		return
	}

	log.Println(chapterNumberInt)

	chapter, err := services.GetChapter(uint(chapterNumberInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()

	queryTelegram := queryParams.Get("telegram")
	queryLocal := queryParams.Get("local")

	if queryTelegram == "true" {
		err = services.SendChapter(*chapter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if queryLocal == "true" {
		var path *string
		for index, page := range chapter.Pages {
			path, err = services.DownloadPage(page.Url, fmt.Sprintf("../chapters/%d_%s", chapter.Number, chapter.Name), fmt.Sprintf("page_%d", index))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		response := SavedDirectory{
			ChapterName:   chapter.Name,
			ChapterNumber: chapter.Number,
			Path:          *path,
		}

		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
}
