package controller

import (
	"net/http"

	svc "github.com/dimaskiddo/frame-go/service"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	svc.ResponseOK(w, "Simple API Go Framework is running")
}

// GetHealth Function to Show Health Check Status
func GetHealth(w http.ResponseWriter, r *http.Request) {
	svc.HealthCheck(w)
}
