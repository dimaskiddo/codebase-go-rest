package controller

import (
	"encoding/json"
	"log"
	"net/http"

	svc "github.com/dimaskiddo/frame-go/service"
)

// GetAuth Function to Get Authorization Token
func GetAuth(w http.ResponseWriter, r *http.Request) {
	var creds svc.BasicCredentials

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&creds)

	// Make Sure Username and Password is Not Empty
	if len(creds.Username) == 0 || len(creds.Password) == 0 {
		svc.ResponseBadRequest(w, "Invalid authorization")
		log.Println("Invalid authorization")
		return
	}

	// Get JWT Token From Pre-Defined Function
	token, err := svc.GetJWTToken(creds.Username)
	if err != nil {
		svc.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	} else {
		jsonToken := map[string]string{"token": token}

		var response svc.FormatGetJWT

		response.Status = true
		response.Code = http.StatusOK
		response.Message = "Success"
		response.Data = jsonToken

		svc.ResponseWrite(w, response.Code, response)
	}
}
