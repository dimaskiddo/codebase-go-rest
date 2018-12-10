package main

import (
	"github.com/dimaskiddo/frame-go/controllers"
	"github.com/dimaskiddo/frame-go/utils"
)

// Routes Initialization Function
func initRoutes() {
	// Set Endpoint for Root Functions
	utils.Router.HandleFunc("/", controllers.GetIndex).Methods("GET")
	utils.Router.HandleFunc("/health", controllers.GetHealth).Methods("GET")

	// Set Endpoint for Upload Functions
	utils.Router.HandleFunc("/uploads", controllers.AddUpload).Methods("POST")

	// Set Endpoint for Authorization Functions
	utils.Router.Handle("/auth", utils.AuthBasic(controllers.GetAuth)).Methods("GET", "POST")

	// Set Endpoint for User Functions
	utils.Router.Handle("/users", utils.AuthJWT(controllers.GetUser)).Methods("GET")
	utils.Router.Handle("/users", utils.AuthJWT(controllers.AddUser)).Methods("POST")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.GetUserByID)).Methods("GET")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.PutUserByID)).Methods("PUT", "PATCH")
	utils.Router.Handle("/users/{id}", utils.AuthJWT(controllers.DelUserByID)).Methods("DELETE")
}
