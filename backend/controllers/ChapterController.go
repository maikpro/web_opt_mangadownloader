package controllers

import (
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
		chapterNameNoSpace := strings.ReplaceAll(chapter.Name, " ", "_")
		zipFileName := fmt.Sprintf("%s.zip", chapterNameNoSpace)

		downloadPath, err := services.DownloadChapter(w, chapter)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("downloaded to %s", *downloadPath)

		// create zip
		buf, err := services.CreateZip(*downloadPath)
		if err != nil {
			log.Fatalf("cannot create zip: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename="+zipFileName)
		w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
		w.Header().Set("Access-Control-Expose-Headers", "Chapter-Name")
		w.Header().Set("Chapter-Name", fmt.Sprintf("%d_%s", chapterNumberInt, chapterNameNoSpace))

		_, err = buf.WriteTo(w)
		if err != nil {
			log.Fatalln("cannot write buffer to writer...")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// w.WriteHeader(http.StatusOK)
}
