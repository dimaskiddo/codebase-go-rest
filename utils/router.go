package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Response Struct
type Response struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Router CORS Configuration Struct
type routerCORSConfig struct {
	Headers []string
	Origins []string
	Methods []string
}

// Router CORS Configuration Variable
var routerCORSCfg routerCORSConfig

// RouterHandler Variable
var RouterHandler http.Handler

// Router Variable
var Router *mux.Router

// InitRouter Function
func initRouter() {
	// Initialize Router
	Router = mux.NewRouter()

	// Set Router Handler with Logging & CORS Support
	RouterHandler = handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedHeaders(routerCORSCfg.Headers),
		handlers.AllowedOrigins(routerCORSCfg.Origins),
		handlers.AllowedMethods(routerCORSCfg.Methods))(Router))
}

// ResponseWrite Function
func ResponseWrite(w http.ResponseWriter, responseCode int, responseData interface{}) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

// ResponseOK Function
func ResponseOK(w http.ResponseWriter, message string) {
	var response Response

	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusOK
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseBadRequest Function
func ResponseBadRequest(w http.ResponseWriter, message string) {
	var response Response

	// Set Default Message
	if len(message) == 0 {
		message = "Bad request"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusBadRequest
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseInternalError Function
func ResponseInternalError(w http.ResponseWriter, message string) {
	var response Response

	// Set Default Message
	if len(message) == 0 {
		message = "Internal server error"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusInternalServerError
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseUnauthorized Function
func ResponseUnauthorized(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "unauthorized"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseAuthenticate Function
func ResponseAuthenticate(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "unauthorized"

	// Write Response
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	w.WriteHeader(response.Code)

	// Write JSON to Response
	json.NewEncoder(w).Encode(response)
}
