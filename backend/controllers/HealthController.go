package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// healthCheckHandler responds to HTTP requests with a simple status message
// @Summary Health check
// @Description Get server health status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func GetHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
