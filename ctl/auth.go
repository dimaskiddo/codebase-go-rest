package ctl

import (
	"encoding/json"
	"net/http"

	"github.com/AasSuhendar/codebase-go-rest/hlp/auth"
	"github.com/AasSuhendar/codebase-go-rest/hlp/router"
)

// GetAuth Function to Get Authorization Token
func GetAuth(w http.ResponseWriter, r *http.Request) {
	var reqBody auth.ReqGetBasic

	// Decode JSON from Request Body to Authorization Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&reqBody)

	// Make Sure Username and Password is Not Empty
	if len(reqBody.Username) == 0 || len(reqBody.Password) == 0 {
		router.ResponseBadRequest(w, "invalid authorization")
		return
	}

	// Get JWT Token From Pre-Defined Function
	token, err := auth.GetJWTToken(reqBody.Username)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	var response auth.ResGetJWT

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data.Token = token

	router.ResponseWrite(w, response.Code, response)
}
