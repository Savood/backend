package image

import (
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/models"
)

// PostUsersIDImageHandler uploads given image and adds a link to the user with the given ID
func PostUsersIDImageHandler (params operations.PostUsersIDImageParams, principal *models.Principal) middleware.Responder {


	return middleware.NotImplemented("")
}

// GetUsersIDImageHandler provides the Image for a User with a given ID
func GetUsersIDImageHandler (params operations.GetUsersIDImageParams, principal *models.Principal) middleware.Responder {


	return middleware.NotImplemented("")
}