package image

import (
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
	"git.dhbw.chd.cx/savood/backend/database/image"
)

// PostUsersIDImageJpegHandler uploads given image
func PostUsersIDImageJpegHandler(params operations.PostUsersIDImageJpegParams, principal *models.Principal) middleware.Responder {

	if params.ID != principal.Userid.String() {
		return operations.NewPostUsersIDImageJpegForbidden()
	}

	img, err := image.ResizeImage(params.Upfile, 0, 0)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	filename := fmt.Sprintf("user_avatar_%s.jpg", params.ID)

	err = image.DeleteImage(filename)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	err = image.UploadImage(filename, img)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDImageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return operations.NewPostUsersIDImageJpegNoContent()

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

// PostUsersIDImageJpegHandler uploads given image
func PostUsersIDBackgroundimageJpegHandler(params operations.PostUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {

	if params.ID != principal.Userid.String() {
		return operations.NewPostUsersIDBackgroundimageJpegForbidden()
	}

	img, err := image.ResizeImage(params.Upfile, 0, 0)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDBackgroundimageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	filename := fmt.Sprintf("user_bg_%s.jpg", params.ID)

	err = image.DeleteImage(filename)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDBackgroundimageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	err = image.UploadImage(filename, img)

	if err != nil {
		str := err.Error()
		return operations.NewGetUsersIDBackgroundimageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return operations.NewPostUsersIDBackgroundimageJpegNoContent()

}

// GetUsersIDImageJpegHandler provides the Image for a User with a given ID
func GetUsersIDBackgroundimageJpegHandler(params operations.GetUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {


	img, err := image.GetImage(fmt.Sprintf("user_bg_%s.jpg", params.ID))
	if err != nil {
		return operations.NewGetUsersIDBackgroundimageJpegNotFound()
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
			return operations.NewGetUsersIDBackgroundimageJpegInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewGetUsersIDBackgroundimageJpegOK().WithPayload(out)
}

