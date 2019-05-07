package service

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// RouterBasePath Variable
var RouterBasePath string

// Router Variable
var Router *chi.Mux

// routerInit Function
func routerInit() {
	// Initialize Router
	Router = chi.NewRouter()

	// Set Router Middleware
	Router.Use(routerCORS)
	Router.Use(routerRealIP)
	Router.Use(routerLogs)
	Router.Use(routerEntitySize)

	// Set Router Handler
	Router.NotFound(handlerNotFound)
	Router.MethodNotAllowed(handlerMethodNotAllowed)
	Router.Get("/favicon.ico", handlerFavIcon)
}

// HealthCheck Function
func HealthCheck(w http.ResponseWriter) {
	// Check Database Connections
	if len(Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := MySQL.Ping()
			if err != nil {
				Log("error", "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		case "mongo":
			err := MongoSession.Ping()
			if err != nil {
				Log("error", "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		}
	}

	// Check Cache Connections
	if len(Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(Config.GetString("CACHE_DRIVER")) {
		case "redis":
			_, err := Redis.Ping().Result()
			if err != nil {
				Log("error", "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		}
	}

	// Return Success
	ResponseSuccess(w, "")
}
