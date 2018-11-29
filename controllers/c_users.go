package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"models"
	"routers"

	"github.com/gorilla/mux"
)

// Function to Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	var response routers.ResponseGetUser

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = models.Users

	// Write Response Data to HTTP
	routers.ResponseWrite(w, response.Code, response)
}

// Function to Get User By ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user []models.User
			var response routers.ResponseGetUser

			// Convert Selected User from Users Array to Single User Array
			user = append(user, models.Users[userID-1])

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"
			response.Data = user

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
	}
}

// Function to Add User
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var response routers.Response

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set User ID to Current Users Array Length + 1
	user.ID = len(models.Users) + 1

	// Insert User to Users Array
	models.Users = append(models.Users, user)

	// Set Response Data
	response.Status = true
	response.Code = http.StatusCreated
	response.Message = "Success"

	// Write Response Data to HTTP
	routers.ResponseWrite(w, response.Code, response)
}

// Function to Update User By ID
func PutUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user models.User
			var response routers.Response

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Update User to Users Array
			models.Users[userID-1].Name = user.Name
			models.Users[userID-1].Email = user.Email

			// Set Response Data
			response.Status = true
			response.Code = http.StatusCreated
			response.Message = "Success"

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
	}
}

// Function to Delete User By ID
func DelUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var response routers.Response

			// Delete User Data from Users Array
			models.Users = append(models.Users[:userID-1], models.Users[userID:]...)

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
	}
}
