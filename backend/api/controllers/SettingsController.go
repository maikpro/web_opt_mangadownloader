package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maikpro/web_opt_mangadownloader/models"
	"github.com/maikpro/web_opt_mangadownloader/services"
)

// GetSettings returns saved Settings.
// @Summary saved settings
// @Description Saved settings
// @Tags settings
// @Accept json
// @Produce json
// @Success 200 {object} models.Settings
// @Router /api/settings [get]
func GetSettings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	settings, err := services.GetSettings()
	if err != nil {
		http.Error(w, "Failed to get settings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings)
}

// SaveSettings saves Settings send through body of a HTTP POST requests.
// @Summary saves settings
// @Description Save your settings
// @Tags settings
// @Accept json
// @Produce json
// @Param settings body models.Settings true "Settings to save"
// @Success 200 {object} models.Settings
// @Router /api/settings [post]
func SaveSettings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var settings models.Settings
	err = json.Unmarshal(body, &settings)
	if err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return
	}

	log.Println("saved data:", settings)

	savedSettings, err := services.SaveSettings(settings)
	if err != nil {
		http.Error(w, "Failed to save settings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedSettings)
}

// UpdateSettings updates Settings send through body of a HTTP PUT requests.
// @Summary updates settings
// @Description Updates your settings
// @Tags settings
// @Accept json
// @Produce json
// @Param settingsId path string true "Updated Settings"
// @Param settings body models.Settings true "Settings to update"
// @Success 200 {object} models.Settings
// @Router /api/settings/id/{settingsId} [put]
func UpdateSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPut {
		log.Println("That's not a PUT Request!")
		http.NotFound(w, r)
		return
	}

	log.Println("UPDATE...")
	//pathVariable := strings.Split(r.URL.Path, "/")
	//settingsId := pathVariable[len(pathVariable)-1]
	settingsId := ps.ByName("id")

	if len(settingsId) == 0 {
		http.Error(w, "Path variable 'id' is required", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var settings models.Settings
	err = json.Unmarshal(body, &settings)
	if err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return
	}

	log.Println(settings.ID)

	if settingsId != settings.ID {
		http.Error(w, "PathVariable ID is not the same like in Body", http.StatusBadRequest)
		return
	}

	updatedSettings, err := services.UpdateSettings(settings)
	if err != nil {
		http.Error(w, "Failed to update settings", http.StatusInternalServerError)
		return
	}

	log.Println("updated data:", updatedSettings)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSettings)
}
