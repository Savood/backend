package image

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
	"git.dhbw.chd.cx/savood/backend/database/image"
	"git.dhbw.chd.cx/savood/backend/dao"
)

// PostOfferingsIDImageHandler uploads given image and adds a link to the offering with the given ID
func PostOfferingsIDImageHandler(params operations.PostOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.ID)
	if err != nil {
		return operations.NewGetOfferingsIDImageInternalServerError()
	}

	if offering.CreatorID != principal.Userid {
		return operations.NewPostOfferingsIDImageForbidden()
	}


	img, err := image.ResizeImage(params.Upfile,0,0)

	if err != nil {
		return operations.NewGetOfferingsIDImageInternalServerError()
	}

	image.UploadImage(fmt.Sprintf("offering_avatar_%s.jpg",params.ID),img)

	return operations.NewPostOfferingsIDImageNoContent()
}

// GetOfferingsIDImageHandler provides the Image for a Offering with a given ID
func GetOfferingsIDImageHandler(params operations.GetOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	img, err := image.GetImage(fmt.Sprintf("offering_avatar_%s.jpg", params.ID))
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

		out, err = image.ResizeImage(img, uint(*params.Width), uint(*params.Height))
		if err != nil {
			str := err.Error()
			return operations.NewGetOfferingsIDImageInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewGetOfferingsIDImageOK().WithPayload(out)
}
