package main

import (
	"github.com/dimaskiddo/frame-go/controllers"
	"github.com/dimaskiddo/frame-go/utils"
)

// Routes Initialization Function
func LoadRoutes() {
	// Initialize Router Endpoint
	utils.Router.HandleFunc("/", controllers.GetIndex).Methods("GET")
	utils.Router.HandleFunc("/uploads", controllers.AddUpload).Methods("POST")

	// Initialize Router Endpoint Secured With Basic Auth
	utils.Router.Handle("/auth", utils.AuthBasic(controllers.GetAuth)).Methods("GET", "POST")

	// Initialize Router Endpoint Secured With Authorization
	utils.Router.Handle("/users", utils.AuthJWT(controllers.GetUser)).Methods("GET")
	utils.Router.Handle("/users", utils.AuthJWT(controllers.AddUser)).Methods("POST")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.GetUserById)).Methods("GET")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.PutUserById)).Methods("PUT", "PATCH")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.DelUserById)).Methods("DELETE")
}
