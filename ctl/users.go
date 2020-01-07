package ctl

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/AasSuhendar/codebase-go-rest/hlp/router"
	"github.com/AasSuhendar/codebase-go-rest/mdl"
)

// resGetUsers Struct
type resGetUsers struct {
	Status  bool       `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []mdl.User `json:"data"`
}

// GetUser Function to Get All User Data
func GetUser(w http.ResponseWriter, r *http.Request) {
	var response resGetUsers

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = mdl.Users

	// Write Response Data to HTTP
	router.ResponseWrite(w, response.Code, response)
}

// AddUser Function to Add User Data
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user mdl.User

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set User ID to Current Users Array Length + 1
	user.ID = len(mdl.Users) + 1

	// Insert User to Users Array
	mdl.Users = append(mdl.Users, user)

	router.ResponseCreated(w)
}

// GetUserByID Function to Get User Data By User ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Check if Requested Data in User Array Range
	if userID <= 0 || userID > len(mdl.Users) {
		router.ResponseBadRequest(w, "invalid array index")
		return
	}

	var users []mdl.User
	var response resGetUsers

	// Convert Selected User from Users Array to Single User Array
	users = append(users, mdl.Users[userID-1])

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Write Response Data to HTTP
	router.ResponseWrite(w, response.Code, response)
}

// PutUserByID Function to Update User Data By User ID
func PutUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Check if Requested Data in User Array Range
	if userID <= 0 || userID > len(mdl.Users) {
		router.ResponseBadRequest(w, "invalid array index")
		return
	}

	var user mdl.User

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Update User to Users Array
	mdl.Users[userID-1].Name = user.Name
	mdl.Users[userID-1].Email = user.Email

	router.ResponseUpdated(w)
}

// DelUserByID Function to Delete User Data By User ID
func DelUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	paramID := chi.URLParam(r, "id")

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		return
	}

	// Check if Requested Data in User Array Range
	if userID <= 0 || userID > len(mdl.Users) {
		router.ResponseBadRequest(w, "invalid array index")
		return
	}

	// Delete User Data from Users Array
	mdl.Users = append(mdl.Users[:userID-1], mdl.Users[userID:]...)

	router.ResponseSuccess(w, "")
}
