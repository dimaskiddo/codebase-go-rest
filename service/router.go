package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// FormatSuccess Struct
type FormatSuccess struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// FormatError Struct
type FormatError struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// Router CORS Configuration Struct
type routerCORSConfig struct {
	Headers []string
	Origins []string
	Methods []string
}

// Router CORS Configuration Variable
var routerCORSCfg routerCORSConfig

// RouterBasePath Variable
var RouterBasePath string

// RouterHandler Variable
var RouterHandler http.Handler

// Router Variable
var Router *mux.Router

// InitRouter Function
func initRouter() {
	// Initialize Router
	Router = mux.NewRouter().StrictSlash(true)

	// Set Router Handler with Logging & CORS Support
	RouterHandler = handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedHeaders(routerCORSCfg.Headers),
		handlers.AllowedOrigins(routerCORSCfg.Origins),
		handlers.AllowedMethods(routerCORSCfg.Methods))(Router))

	// Set Router Default Not Found Handler
	Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ResponseNotFound(w)
	})
}

// HealthCheck Function
func HealthCheck(w http.ResponseWriter) {
	// Check Any Database Connections
	if len(Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := MySQL.Ping()
			if err != nil {
				ResponseInternalError(w, "Database MySQL Connection Failed")
				log.Fatalln("Database MySQL Connection Failed")
				return
			}
		case "mongo":
			err := MongoSession.Ping()
			if err != nil {
				ResponseInternalError(w, "Database Mongo Connection Failed")
				log.Fatalln("Database Mongo Connection Failed")
				return
			}
		}
	}

	// Check Any Cache Connections
	if len(Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(Config.GetString("CACHE_DRIVER")) {
		case "redis":
			_, err := Redis.Ping().Result()
			if err != nil {
				ResponseInternalError(w, "Cache Redis Connection Failed")
				log.Fatalln("Cache Redis Connection Failed")
				return
			}
		}
	}

	// Return OK
	ResponseSuccess(w, "")
}

// ResponseWrite Function
func ResponseWrite(w http.ResponseWriter, responseCode int, responseData interface{}) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

// ResponseSuccess Function
func ResponseSuccess(w http.ResponseWriter, message string) {
	var response FormatSuccess

	// Set Default Message
	if len(message) == 0 {
		message = "Success"
	}

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseNotFound Function
func ResponseNotFound(w http.ResponseWriter) {
	var response FormatError

	// Set Response Data
	response.Status = false
	response.Code = http.StatusNotFound
	response.Message = "Not Found"
	response.Error = "Not Found"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseBadRequest Function
func ResponseBadRequest(w http.ResponseWriter, message string) {
	var response FormatError

	// Set Default Message
	if len(message) == 0 {
		message = "Bad Request"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusBadRequest
	response.Message = "Bad Request"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseInternalError Function
func ResponseInternalError(w http.ResponseWriter, message string) {
	var response FormatError

	// Set Default Message
	if len(message) == 0 {
		message = "Internal Server Error"
	}

	// Set Response Data
	response.Status = false
	response.Code = http.StatusInternalServerError
	response.Message = "Internal Server Error"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseUnauthorized Function
func ResponseUnauthorized(w http.ResponseWriter) {
	var response FormatError

	// Set Response Data
	response.Status = false
	response.Code = http.StatusUnauthorized
	response.Message = "Unauthorized"
	response.Error = "Unaothorized"

	// Set Response Data to HTTP
	ResponseWrite(w, response.Code, response)
}

// ResponseAuthenticate Function
func ResponseAuthenticate(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
	ResponseUnauthorized(w)
}
