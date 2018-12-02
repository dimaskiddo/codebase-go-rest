package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dimaskiddo/frame-go/models"
	"github.com/dimaskiddo/frame-go/utils"

	"github.com/gorilla/mux"
)

// Get User Response Structure
type ResponseGetUser struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

// Function to Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	var response ResponseGetUser

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "success"
	response.Data = models.Users

	// Write Response Data to HTTP
	utils.ResponseWrite(w, response.Code, response)
}

// Function to Add User
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var response utils.Response

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set User ID to Current Users Array Length + 1
	user.ID = len(models.Users) + 1

	// Insert User to Users Array
	models.Users = append(models.Users, user)

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "success"

	// Write Response Data to HTTP
	utils.ResponseWrite(w, response.Code, response)
}

// Function to Get User By ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
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
			response.Message = "success"
			response.Data = users

			// Write Response Data to HTTP
			utils.ResponseWrite(w, response.Code, response)
		} else {
			// Lower Case Error Message
			errMessage := strings.ToLower("invalid array index")

			// Write Error Response Data to HTTP
			utils.ResponseBadRequest(w, errMessage)

			// Print Log Error to Router Log Handler
			log.Println(errMessage)
		}
	} else {
		// Lower Case Error Message
		errMessage := strings.ToLower(err.Error())

		// Write Error Response Data to HTTP
		utils.ResponseInternalError(w, errMessage)

		// Print Log Error to Router Log Handler
		log.Println(errMessage)
	}
}

// Function to Update User By ID
func PutUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(models.Users) > 0 && userID <= len(models.Users) {
			var user models.User
			var response utils.Response

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Update User to Users Array
			models.Users[userID-1].Name = user.Name
			models.Users[userID-1].Email = user.Email

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "success"

			// Write Response Data to HTTP
			utils.ResponseWrite(w, response.Code, response)
		} else {
			// Lower Case Error Message
			errMessage := strings.ToLower("invalid array index")

			// Write Error Response Data to HTTP
			utils.ResponseBadRequest(w, errMessage)

			// Print Log Error to Router Log Handler
			log.Println(errMessage)
		}
	} else {
		// Lower Case Error Message
		errMessage := strings.ToLower(err.Error())

		// Write Error Response Data to HTTP
		utils.ResponseInternalError(w, errMessage)

		// Print Log Error to Router Log Handler
		log.Println(errMessage)
	}
}

// Function to Delete User By ID
func DelUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Get ID Parameters From URI Then Convert it to Integer
	userID, err := strconv.Atoi(params["id"])
	if err == nil {
		// Check if Requested Data in User Array Range
		if len(models.Users) > 0 && userID <= len(models.Users) {
			var response utils.Response

			// Delete User Data from Users Array
			models.Users = append(models.Users[:userID-1], models.Users[userID:]...)

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "success"

			// Write Response Data to HTTP
			utils.ResponseWrite(w, response.Code, response)
		} else {
			// Lower Case Error Message
			errMessage := strings.ToLower("invalid array index")

			// Write Error Response Data to HTTP
			utils.ResponseBadRequest(w, errMessage)

			// Print Log Error to Router Log Handler
			log.Println(errMessage)
		}
	} else {
		// Lower Case Error Message
		errMessage := strings.ToLower(err.Error())

		// Write Error Response Data to HTTP
		utils.ResponseInternalError(w, errMessage)

		// Print Log Error to Router Log Handler
		log.Println(errMessage)
	}
}
