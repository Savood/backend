package image

import (
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
	"git.dhbw.chd.cx/savood/backend/database/image"
)

// PostUsersIDImageJpegHandler uploads given image and adds a link to the user with the given ID
func PostUsersIDImageJpegHandler(params operations.PostUsersIDImageJpegParams, principal *models.Principal) middleware.Responder {


	return middleware.NotImplemented("")
}

// GetUsersIDImageJpegHandler provides the Image for a User with a given ID
func GetUsersIDImageJpegHandler(params operations.GetUsersIDImageJpegParams, principal *models.Principal) middleware.Responder {


	img, err := image.GetImage(fmt.Sprintf("user_avatar_%s.jpg", params.ID))
	if err != nil {
		return operations.NewGetUsersIDImageJpegNotFound()
	}

	out := img

	heightDefault := float64(0)
	widthDefault := float64(0)

	if params.Height != nil || params.Width != nil {
		if params.Height == nil || *params.Height < float64(0) {
			params.Height = &heightDefault
		}
		if params.Width == nil || *params.Width < float64(0) {
			params.Width = &widthDefault
		}

		out, err = image.ResizeImage(img, uint(*params.Width), uint(*params.Height))
		if err != nil {
			str := err.Error()
			return operations.NewGetUsersIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewGetUsersIDImageJpegOK().WithPayload(out)
}