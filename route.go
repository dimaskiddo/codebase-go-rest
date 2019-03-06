package main

import (
	ctl "github.com/dimaskiddo/frame-go/controller"
	svc "github.com/dimaskiddo/frame-go/service"
)

// Routes Initialization Function
func initRoutes() {
	// Set Endpoint for Root Functions
	svc.Router.HandleFunc(svc.RouterBasePath+"/", ctl.GetIndex).Methods("GET")
	svc.Router.HandleFunc(svc.RouterBasePath+"/health", ctl.GetHealth).Methods("GET")

	// Set Endpoint for Authorization Functions
	svc.Router.Handle(svc.RouterBasePath+"/auth", svc.AuthBasic(ctl.GetAuth)).Methods("GET", "POST")

	// Set Endpoint for User Functions
	svc.Router.Handle(svc.RouterBasePath+"/users", svc.AuthJWT(ctl.GetUser)).Methods("GET")
	svc.Router.Handle(svc.RouterBasePath+"/users", svc.AuthJWT(ctl.AddUser)).Methods("POST")
	svc.Router.Handle(svc.RouterBasePath+"/users/{id}", svc.AuthJWT(ctl.GetUserByID)).Methods("GET")
	svc.Router.Handle(svc.RouterBasePath+"/users/{id}", svc.AuthJWT(ctl.PutUserByID)).Methods("PUT", "PATCH")
	svc.Router.Handle(svc.RouterBasePath+"/users/{id}", svc.AuthJWT(ctl.DelUserByID)).Methods("DELETE")

	// Set Endpoint for Upload Function
	svc.Router.Handle(svc.RouterBasePath+"/upload", svc.AuthJWT(ctl.UploadFile)).Methods("POST")
}
