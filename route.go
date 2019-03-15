package main

import (
	ctl "github.com/dimaskiddo/codebase-go-rest/controller"
	svc "github.com/dimaskiddo/codebase-go-rest/service"
)

// RoutesInit Function
func routesInit() {
	// Set Endpoint for Root Functions
	svc.Router.Get(svc.RouterBasePath+"/", ctl.GetIndex)
	svc.Router.Get(svc.RouterBasePath+"/health", ctl.GetHealth)

	// Set Endpoint for Authorization Functions
	svc.Router.With(svc.AuthBasic).Get(svc.RouterBasePath+"/auth", ctl.GetAuth)

	// Set Endpoint for User Functions
	svc.Router.With(svc.AuthJWT).Get(svc.RouterBasePath+"/users", ctl.GetUser)
	svc.Router.With(svc.AuthJWT).Post(svc.RouterBasePath+"/users", ctl.AddUser)
	svc.Router.With(svc.AuthJWT).Get(svc.RouterBasePath+"/users/{id}", ctl.GetUserByID)
	svc.Router.With(svc.AuthJWT).Put(svc.RouterBasePath+"/users/{id}", ctl.PutUserByID)
	svc.Router.With(svc.AuthJWT).Patch(svc.RouterBasePath+"/users/{id}", ctl.PutUserByID)
	svc.Router.With(svc.AuthJWT).Delete(svc.RouterBasePath+"/users/{id}", ctl.DelUserByID)

	// Set Endpoint for Upload Function
	svc.Router.With(svc.AuthJWT).Post(svc.RouterBasePath+"/upload", ctl.UploadFile)
}
