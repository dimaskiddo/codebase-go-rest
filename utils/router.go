package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Default Response Structure
type Response struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Router CORS Configuration Struct
type RouterCORSConfiguration struct {
	Headers []string
	Origins []string
	Methods []string
}

// Router CORS Configuration Variable
var RouterCORS RouterCORSConfiguration

// Router Handler Variable
var RouterHandler http.Handler

// Router Variable
var Router *mux.Router

// Router Initialize Function
func InitRouter() {
	// Initialize Router
	Router = mux.NewRouter()

	// Set Router Handler with Logging & CORS Support
	RouterHandler = handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedHeaders(RouterCORS.Headers),
		handlers.AllowedOrigins(RouterCORS.Origins),
		handlers.AllowedMethods(RouterCORS.Methods))(Router))
}

// Write Response to HTTP Writer
func ResponseWrite(w http.ResponseWriter, responseCode int, responseData interface{}) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

// Write Response Bad Request
func ResponseBadRequest(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusBadRequest
	response.Message = "Bad Request"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Internal Server Error
func ResponseInternalError(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusInternalServerError
	response.Message = "Internal Server Error"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Unauthorized
func ResponseUnauthorized(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "Unauthorized"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Authenticate
func ResponseAuthenticate(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "Unauthorized"

	// Write Response
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(response.Code)

	// Write JSON to Response
	json.NewEncoder(w).Encode(response)
}