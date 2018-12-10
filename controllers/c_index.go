package controllers

import (
	"net/http"

	"github.com/dimaskiddo/frame-go/utils"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	utils.ResponseOK(w, "Simple API Go Framework is running")
}

// GetHealth Function to Show Health Check Status
func GetHealth(w http.ResponseWriter, r *http.Request) {
	utils.HealthCheck(w)
}
