package controllers

import (
	"net/http"

	"github.com/dimaskiddo/frame-go/utils"
)

// Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var response utils.Response

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Simple Go Framework is running"

	utils.ResponseWrite(w, response.Code, response)
}
