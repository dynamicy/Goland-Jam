package controllers

import (
	"net/http"
)

// HealthCheckHandler handles the health check request
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
