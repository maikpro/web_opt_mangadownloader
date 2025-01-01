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

type IOPTClient interface {
	GetChapter(chapterNumber uint) (*models.Chapter, error)
	DownloadChapter(chapter *models.Chapter) (*string, error)
	GetArcList() ([]models.Arc, error)
}

type OPTClient struct{}

type ChapterWrapper struct {
	Chapter models.Chapter `json:"chapter"`
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

func (optClient *OPTClient) GetChapter(chapterNumber uint) (*models.Chapter, error) {
	// note: One-Piece-Tube does not provide Manga from Chapter 1 - 419
	if chapterNumber < 420 {
		err := errors.New("OPTClient: onepiece-tube.com does not provide Manga from Chapter 1 - 419")
		log.Fatal(err.Error())
		return nil, err
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
	// scriptRawText: window.__data = {"chapter":{"name":"Keine Spur von Zorro","pages":[{"url":"https:\/\/onepiece.tube\/upload\/manga\/kapitel\/0512-515\/01.jpg","height":1200,"width":826,"type":"image\/jpeg"}, ...
	parsedJsonString := strings.Replace(strings.Split(scriptRawText, "=")[1], ";", "", -1)

	var chapterWrapper ChapterWrapper
	err = json.Unmarshal([]byte(parsedJsonString), &chapterWrapper)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	chapter := &chapterWrapper.Chapter
	chapter.Number = chapterNumber

	return chapter, nil
}

func (optClient *OPTClient) DownloadChapter(chapter *models.Chapter) (*string, error) {
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

func getPageImage(url string) ([]byte, error) {
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

func downloadPage(url string, savePath string, filename string) (*string, error) {
	imageData, err := getPageImage(url)
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

func (optClient *OPTClient) GetArcList() ([]models.Arc, error) {
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
