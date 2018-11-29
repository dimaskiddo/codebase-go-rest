package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/frame-go/controllers"
	"github.com/dimaskiddo/frame-go/dbs"
	"github.com/dimaskiddo/frame-go/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Channel for OS Signal
	signalOS := make(chan os.Signal, 1)

	// Initialize Configuration
	utils.ConfigInitialize()

	// Initialize Database
	dbs.DatabaseInitialize()

	// Initialize Router
	router := mux.NewRouter()

	// Initialize Router Endpoint
	router.HandleFunc("/", controllers.GetIndex).Methods("GET")

	// Initialize Router Endpoint Secured With Basic Auth
	router.Handle("/auth", utils.AuthBasic(controllers.GetAuthentication)).Methods("GET", "POST")

	// Initialize Router Endpoint Secured With Authorization
	router.Handle("/users", utils.AuthJWT(controllers.GetUser)).Methods("GET")
	router.Handle("/users", utils.AuthJWT(controllers.AddUser)).Methods("POST")
	router.Handle("/users/{id}", utils.AuthJWT(controllers.GetUserById)).Methods("GET")
	router.Handle("/users/{id}", utils.AuthJWT(controllers.PutUserById)).Methods("PUT", "PATCH")
	router.Handle("/users/{id}", utils.AuthJWT(controllers.DelUserById)).Methods("DELETE")

	// Set Router Handler with Logging & CORS Support
	routerHandler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedHeaders(utils.RouterCORS.Headers),
		handlers.AllowedOrigins(utils.RouterCORS.Origins),
		handlers.AllowedMethods(utils.RouterCORS.Methods))(router))

	// Initialize Server With Initialized Router
	server := utils.NewServer(routerHandler)

	// Starting Server
	server.Start()

	// Catch OS Signal from Channel
	signal.Notify(signalOS, os.Interrupt, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-signalOS

	// Add Some Spaces When Done
	fmt.Println("Stopping Service")

	// Defer Some Function Before End
	defer server.Stop()
	switch strings.ToLower(utils.Config.GetString("DB_DRIVER")) {
	case "mysql":
		defer dbs.MySQL.Close()
	case "mongo":
		defer dbs.MongoSession.Close()
	}
}
