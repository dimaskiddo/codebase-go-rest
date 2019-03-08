package controller

import (
	"encoding/json"
	"net/http"

	svc "github.com/dimaskiddo/codebase-go-rest/service"
)

// GetAuth Function to Get Authorization Token
func GetAuth(w http.ResponseWriter, r *http.Request) {
	var reqBody svc.ReqGetBasic

	// Decode JSON from Request Body to Authorization Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&reqBody)

	// Make Sure Username and Password is Not Empty
	if len(reqBody.Username) == 0 || len(reqBody.Password) == 0 {
		svc.ResponseBadRequest(w, "invalid authorization")
		return
	}

	// Get JWT Token From Pre-Defined Function
	token, err := svc.GetJWTToken(reqBody.Username)
	if err != nil {
		svc.ResponseInternalError(w, err.Error())
		return
	}

	var response svc.ResGetJWT

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data.Token = token

	svc.ResponseWrite(w, response.Code, response)
}
