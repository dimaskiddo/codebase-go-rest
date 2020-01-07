package main

import (
	"github.com/AasSuhendar/codebase-go-rest/ctl"
	"github.com/AasSuhendar/codebase-go-rest/hlp/auth"
	"github.com/AasSuhendar/codebase-go-rest/hlp/router"
)

// Initialize Function in Main Route
func init() {
	// Set Endpoint for Root Functions
	router.Router.Get(router.RouterBasePath+"/", ctl.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", ctl.GetHealth)

	// Set Endpoint for Authorization Functions
	router.Router.With(auth.Basic).Get(router.RouterBasePath+"/auth", ctl.GetAuth)

	// Set Endpoint for User Functions
	router.Router.With(auth.JWT).Get(router.RouterBasePath+"/users", ctl.GetUser)
	router.Router.With(auth.JWT).Post(router.RouterBasePath+"/users", ctl.AddUser)
	router.Router.With(auth.JWT).Get(router.RouterBasePath+"/users/{id}", ctl.GetUserByID)
	router.Router.With(auth.JWT).Put(router.RouterBasePath+"/users/{id}", ctl.PutUserByID)
	router.Router.With(auth.JWT).Patch(router.RouterBasePath+"/users/{id}", ctl.PutUserByID)
	router.Router.With(auth.JWT).Delete(router.RouterBasePath+"/users/{id}", ctl.DelUserByID)

	// Set Endpoint for Upload Function
	router.Router.With(auth.JWT).Post(router.RouterBasePath+"/upload", ctl.UploadFile)
}
