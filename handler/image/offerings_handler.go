package image

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
	"git.dhbw.chd.cx/savood/backend/database/image"
	"git.dhbw.chd.cx/savood/backend/dao"
)

// PostOfferingsIDImageJpegHandler uploads given image
func PostOfferingsIDImageJpegHandler(params operations.PostOfferingsIDImageJpegParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.ID, nil)
	if err != nil {
		str := err.Error()
		return operations.NewGetOfferingsIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.Creator.ID != principal.Userid {
		return operations.NewPostOfferingsIDImageJpegForbidden()
	}

	img, err := image.ResizeImage(params.Upfile, 0, 0)

	if err != nil {
		str := err.Error()
		return operations.NewGetOfferingsIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	filename := fmt.Sprintf("offering_avatar_%s.jpg", params.ID)

	image.DeleteImage(filename)

	err = image.UploadImage(filename, img)

	if err != nil {
		str := err.Error()
		return operations.NewPostOfferingsIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return operations.NewPostOfferingsIDImageJpegNoContent()
}

// GetOfferingsIDImageJpegHandler provides the Image for a Offering with a given ID
func GetOfferingsIDImageJpegHandler(params operations.GetOfferingsIDImageJpegParams, principal *models.Principal) middleware.Responder {

	img, err := image.GetImage(fmt.Sprintf("offering_avatar_%s.jpg", params.ID))
	if err != nil {
		return operations.NewGetOfferingsIDImageJpegNotFound()
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
			return operations.NewGetOfferingsIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewGetOfferingsIDImageJpegOK().WithPayload(out)
}
