package controllers

import (
	"encoding/json"
	"net/http"

	"routers"
	"utils"
)

// Function to Get Authentication Using JWT
func GetAuthentication(w http.ResponseWriter, r *http.Request) {
	var creds utils.JWTCredentials

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&creds)

	// Make Sure Username and Password is Not Empty
	if len(creds.Username) == 0 || len(creds.Password) == 0 {
		routers.ResponseBadRequest(w)
		return
	}

	// Some Business Logic Here to Match The Username and Password
	if creds.Username == "user" && creds.Password == "password" {
		// Get JWT Token From Pre-Defined Function
		token, err := utils.GetJWTToken(creds.Username)
		if err != nil {
			routers.ResponseInternalError(w)
		} else {
			var response utils.JWTResponse

			response.Status = true
			response.Code = http.StatusOK
			response.Token = token

			routers.ResponseWrite(w, response.Code, response)
		}
	} else {
		routers.ResponseUnauthorized(w)
	}
}
