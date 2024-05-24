package services

import (
	"log"

	"github.com/maikpro/web_opt_mangadownloader/database"
	"github.com/maikpro/web_opt_mangadownloader/models"
)

func GetSettings() (*models.Settings, error) {
	settings, err := database.GetSettings()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return settings, nil
}

func SaveSettings(settings models.Settings) (*models.Settings, error) {
	settingsSaved, err := database.SaveSettings(settings)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return settingsSaved, nil
}

func UpdateSettings(settings models.Settings) (*models.Settings, error) {
	settingsUpdated, err := database.UpdateSettings(settings)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return settingsUpdated, nil
}
