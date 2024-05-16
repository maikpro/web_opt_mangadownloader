package services

import (
	"log"

	"github.com/maikpro/web_opt_mangadownloader/database"
	"github.com/maikpro/web_opt_mangadownloader/models"
)

func SaveSettings(settings models.Settings) (*models.Settings, error) {
	settingsSaved, err := database.Save(settings)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return settingsSaved, nil
}
