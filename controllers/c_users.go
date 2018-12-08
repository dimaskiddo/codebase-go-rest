package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dimaskiddo/frame-go/models"
	"github.com/dimaskiddo/frame-go/utils"

	"github.com/gorilla/mux"
)

// ResponseGetUser Struct
type ResponseGetUser struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

// GetUser Function to Get All User Data
func GetUser(w http.ResponseWriter, r *http.Request) {
	var response ResponseGetUser

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = models.Users

	// Write Response Data to HTTP
	utils.ResponseWrite(w, response.Code, response)
}

// AddUser Function to Add User Data
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set User ID to Current Users Array Length + 1
	user.ID = len(models.Users) + 1

	// Insert User to Users Array
	models.Users = append(models.Users, user)

	utils.ResponseOK(w, "")
}

// GetUserByID Function to Get User Data By User ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(models.Users) > 0 && userID <= len(models.Users) {
			var users []models.User
			var response ResponseGetUser

			// Convert Selected User from Users Array to Single User Array
			users = append(users, models.Users[userID-1])

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"
			response.Data = users

			// Write Response Data to HTTP
			utils.ResponseWrite(w, response.Code, response)
		} else {
			utils.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		utils.ResponseInternalError(w, err.Error())
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
		if len(models.Users) > 0 && userID <= len(models.Users) {
			var user models.User

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Update User to Users Array
			models.Users[userID-1].Name = user.Name
			models.Users[userID-1].Email = user.Email

			utils.ResponseOK(w, "")
		} else {
			utils.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		utils.ResponseInternalError(w, err.Error())
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
		if len(models.Users) > 0 && userID <= len(models.Users) {
			// Delete User Data from Users Array
			models.Users = append(models.Users[:userID-1], models.Users[userID:]...)

			utils.ResponseOK(w, "")
		} else {
			utils.ResponseBadRequest(w, "Invalid array index")
			log.Println("Invalid array index")
		}
	} else {
		utils.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}
