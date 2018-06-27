package image

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
)

// PostOfferingsIDImageHandler uploads given image and adds a link to the offering with the given ID
func PostOfferingsIDImageHandler(params operations.PostOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	//uploadImage("test", params.Upfile)

	return middleware.NotImplemented("")
}

// GetOfferingsIDImageHandler provides the Image for a Offering with a given ID
func GetOfferingsIDImageHandler(params operations.GetOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	img, err := getImage(fmt.Sprintf("offering_avatar_%s.jpg", params.ID))
	if err != nil {
		return operations.NewGetOfferingsIDImageNotFound()
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

		out, err = resizeImage(img, uint(*params.Width), uint(*params.Height))
		if err != nil {
			str := err.Error()
			return operations.NewGetOfferingsIDImageInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewGetOfferingsIDImageOK().WithPayload(out)
}
