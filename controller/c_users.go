package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	mdl "github.com/dimaskiddo/frame-go/model"
	svc "github.com/dimaskiddo/frame-go/service"

	"github.com/gorilla/mux"
)

// ResponseGetUser Struct
type ResponseGetUser struct {
	Status  bool       `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []mdl.User `json:"data"`
}

// GetUser Function to Get All User Data
func GetUser(w http.ResponseWriter, r *http.Request) {
	var response ResponseGetUser

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = mdl.Users

	// Write Response Data to HTTP
	svc.ResponseWrite(w, response.Code, response)
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

	svc.ResponseOK(w, "")
}

// GetUserByID Function to Get User Data By User ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(mdl.Users) > 0 && userID <= len(mdl.Users) {
			var users []mdl.User
			var response ResponseGetUser

			// Convert Selected User from Users Array to Single User Array
			users = append(users, mdl.Users[userID-1])

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"
			response.Data = users

			// Write Response Data to HTTP
			svc.ResponseWrite(w, response.Code, response)
		} else {
			svc.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		svc.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}

// PutUserByID Function to Update User Data By User ID
func PutUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(mdl.Users) > 0 && userID <= len(mdl.Users) {
			var user mdl.User

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Update User to Users Array
			mdl.Users[userID-1].Name = user.Name
			mdl.Users[userID-1].Email = user.Email

			svc.ResponseOK(w, "")
		} else {
			svc.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		svc.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}

// DelUserByID Function to Delete User Data By User ID
func DelUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(mdl.Users) > 0 && userID <= len(mdl.Users) {
			// Delete User Data from Users Array
			mdl.Users = append(mdl.Users[:userID-1], mdl.Users[userID:]...)

			svc.ResponseOK(w, "")
		} else {
			svc.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		svc.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}
