package main

import (
	ctl "github.com/dimaskiddo/frame-go/controller"
	svc "github.com/dimaskiddo/frame-go/service"
)

// Routes Initialization Function
func initRoutes() {
	// Set Endpoint for Root Functions
	svc.Router.HandleFunc("/", ctl.GetIndex).Methods("GET")
	svc.Router.HandleFunc("/health", ctl.GetHealth).Methods("GET")

	// Set Endpoint for Upload Functions
	svc.Router.HandleFunc("/uploads", ctl.AddUpload).Methods("POST")

	// Set Endpoint for Authorization Functions
	svc.Router.Handle("/auth", svc.AuthBasic(ctl.GetAuth)).Methods("GET", "POST")

	// Set Endpoint for User Functions
	svc.Router.Handle("/users", svc.AuthJWT(ctl.GetUser)).Methods("GET")
	svc.Router.Handle("/users", svc.AuthJWT(ctl.AddUser)).Methods("POST")
	svc.Router.Handle("/users/{id}", svc.AuthJWT(ctl.GetUserByID)).Methods("GET")
	svc.Router.Handle("/users/{id}", svc.AuthJWT(ctl.PutUserByID)).Methods("PUT", "PATCH")
	svc.Router.Handle("/users/{id}", svc.AuthJWT(ctl.DelUserByID)).Methods("DELETE")
}
