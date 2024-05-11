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

type ResponseData struct {
	Path string `json:"path"`
}

func DownloadChapter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("That's not a POST Request!")
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

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

	var path *string
	for index, page := range chapter.Pages {
		path, err = services.DownloadPage(page.Url, fmt.Sprintf("../chapters/%d_%s", chapter.Number, chapter.Name), fmt.Sprintf("page_%d", index))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	response := ResponseData{
		Path: *path,
	}

	json.NewEncoder(w).Encode(response)
}
