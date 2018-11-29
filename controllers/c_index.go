package controllers

import (
	"net/http"

	"github.com/dimaskiddo/frame-go/routers"
)

// Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var response routers.Response

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Simple Go Programming Example (API Simple)"

	routers.ResponseWrite(w, response.Code, response)
}
