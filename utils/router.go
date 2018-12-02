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
type routerCORSConfig struct {
	Headers []string
	Origins []string
	Methods []string
}

// Router CORS Configuration Variable
var routerCORSCfg routerCORSConfig

// Router Handler Variable
var RouterHandler http.Handler

// Router Variable
var Router *mux.Router

// Router Initialize Function
func initRouter() {
	// Initialize Router
	Router = mux.NewRouter()

	// Set Router Handler with Logging & CORS Support
	RouterHandler = handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedHeaders(routerCORSCfg.Headers),
		handlers.AllowedOrigins(routerCORSCfg.Origins),
		handlers.AllowedMethods(routerCORSCfg.Methods))(Router))
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
func ResponseBadRequest(w http.ResponseWriter, msg string) {
	var response Response

	// Set Default Message
	if len(msg) == 0 {
		msg = "bad request"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusBadRequest
	response.Message = msg

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Internal Server Error
func ResponseInternalError(w http.ResponseWriter, msg string) {
	var response Response

	// Set Default Message
	if len(msg) == 0 {
		msg = "internal server error"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusInternalServerError
	response.Message = msg

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Unauthorized
func ResponseUnauthorized(w http.ResponseWriter) {
	var response Response

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "unauthorized"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// Write Response Authenticate
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
