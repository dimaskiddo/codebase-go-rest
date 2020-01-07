package router

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/AasSuhendar/codebase-go-rest/hlp"
	"github.com/AasSuhendar/codebase-go-rest/hlp/cache"
	"github.com/AasSuhendar/codebase-go-rest/hlp/db"
)

// RouterBasePath Variable
var RouterBasePath string

// Router Variable
var Router *chi.Mux

// Initialize Function in Router
func init() {
	// Initialize Router
	Router = chi.NewRouter()
	RouterBasePath = hlp.Config.GetString("ROUTER_BASE_PATH")

	// Set Router CORS Configuration
	routerCORSCfg.Origins = hlp.Config.GetString("CORS_ALLOWED_ORIGIN")
	routerCORSCfg.Methods = hlp.Config.GetString("CORS_ALLOWED_METHOD")
	routerCORSCfg.Headers = hlp.Config.GetString("CORS_ALLOWED_HEADER")

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
	if len(hlp.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(hlp.Config.GetString("DB_DRIVER")) {
		case "mysql":
			err := db.MySQL.Ping()
			if err != nil {
				hlp.LogPrintln(hlp.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		case "mongo":
			err := db.MongoSession.Ping()
			if err != nil {
				hlp.LogPrintln(hlp.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		}
	}

	// Check Cache Connections
	if len(hlp.Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(hlp.Config.GetString("CACHE_DRIVER")) {
		case "redis":
			_, err := cache.Redis.Ping().Result()
			if err != nil {
				hlp.LogPrintln(hlp.LogLevelError, "health-check", err.Error())
				ResponseInternalError(w, err.Error())
				return
			}
		}
	}

	// Return Success
	ResponseSuccess(w, "")
}
