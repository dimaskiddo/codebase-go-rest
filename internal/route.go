package internal

import (
	"github.com/dimaskiddo/codebase-go-rest/pkg/auth"
	"github.com/dimaskiddo/codebase-go-rest/pkg/router"

	"github.com/dimaskiddo/codebase-go-rest/internal/index"
	"github.com/dimaskiddo/codebase-go-rest/internal/uploads"
	"github.com/dimaskiddo/codebase-go-rest/internal/users"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {
	// Set Endpoint for Root Functions
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	// Set Endpoint for Authorization Functions
	router.Router.With(auth.Basic).Get(router.RouterBasePath+"/auth", index.GetAuth)

	// Set Endpoint for User Functions
	router.Router.With(auth.JWT).Get(router.RouterBasePath+"/users", users.GetUser)
	router.Router.With(auth.JWT).Post(router.RouterBasePath+"/users", users.AddUser)
	router.Router.With(auth.JWT).Get(router.RouterBasePath+"/users/{id}", users.GetUserByID)
	router.Router.With(auth.JWT).Put(router.RouterBasePath+"/users/{id}", users.PutUserByID)
	router.Router.With(auth.JWT).Patch(router.RouterBasePath+"/users/{id}", users.PutUserByID)
	router.Router.With(auth.JWT).Delete(router.RouterBasePath+"/users/{id}", users.DelUserByID)

	// Set Endpoint for Upload Function
	router.Router.With(auth.JWT).Post(router.RouterBasePath+"/uploads", uploads.UploadFile)
}
