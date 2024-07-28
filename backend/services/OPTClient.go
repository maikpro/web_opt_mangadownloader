package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/maikpro/web_opt_mangadownloader/models"
)

func GetChapter(chapterNumber uint) (*models.Chapter, error) {
	// note: One-Piece-Tube does not provide Manga from Chapter 1 - 419
	if chapterNumber < 420 {
		return nil, errors.New("onepiece-tube.com does not provide Manga from Chapter 1 - 419")
	}

	url := fmt.Sprintf("https://onepiece-tube.com/manga/kapitel/%d", chapterNumber)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scriptRawText := doc.Find("#app > script").First().Text()
	jsonString := strings.Replace(strings.Split(scriptRawText, "=")[1], ";", "", -1)
	var newjson models.Data
	json.Unmarshal([]byte(jsonString), &newjson)
	chapter := &newjson.Chapter
	chapter.Number = chapterNumber
	return chapter, nil
}

func GetPageImage(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return nil, err
	}
	defer response.Body.Close()
	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

func DownloadChapter(w http.ResponseWriter, chapter *models.Chapter) (*string, error) {
	var downloadPath *string
	var err error

	for index, page := range chapter.Pages {
		filePath := fmt.Sprintf("../chapters/%d_%s", chapter.Number, chapter.Name)
		modifiedFilePath := strings.ReplaceAll(filePath, " ", "_")
		downloadPath, err = downloadPage(page.Url, modifiedFilePath, fmt.Sprintf("page_%d", index))
		if err != nil {
			return nil, err
		}
	}

	return downloadPath, nil
}

func downloadPage(url string, savePath string, filename string) (*string, error) {
	imageData, err := GetPageImage(url)
	if err != nil {
		return nil, err
	}

	// Get the extension from image
	// Get the file extension from the URL
	ext := filepath.Ext(url)

	// Remove the leading dot from the extension
	ext = ext[1:]
	filename = fmt.Sprintf("%s.%s", filename, ext)
	folderPath, err := SaveFile(savePath, filename, imageData)
	if err != nil {
		log.Println("Error saving image locally:", err)
		return nil, err
	}

	return &folderPath, err
}

type OPTListData struct {
	Arcs    []OPTArc          `json:"arcs"`
	Entries []models.OPTEntry `json:"entries"`
}

// {"id":41,"name":"Egghead Arc","min":1058,"max":1113}
type OPTArc struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

func GetArcList() ([]models.Arc, error) {
	res, err := http.Get("https://onepiece-tube.com/manga/kapitel-mangaliste")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scriptRawText := doc.Find("#main-content > script").First().Text()
	jsonString := strings.Replace(strings.Split(scriptRawText, "=")[1], ";", "", -1)
	var optListData OPTListData
	err = json.Unmarshal([]byte(jsonString), &optListData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	arcs := mapIntoArcs(optListData)
	return arcs, nil
}

func mapIntoArcs(optListData OPTListData) []models.Arc {
	var arcs []models.Arc

	for _, optListArc := range optListData.Arcs {
		var arc models.Arc
		arc.Name = strings.TrimRight(optListArc.Name, " ")

		for _, optEntry := range optListData.Entries {
			min := optListArc.Min
			max := optListArc.Max

			if optEntry.Number <= max && optEntry.Number >= min {
				arc.Entries = append(arc.Entries, optEntry)
			}
		}

		arcs = append(arcs, arc)
	}
	return arcs
}
