package image

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
)

// PostOfferingsIDImageHandler uploads given image and adds a link to the offering with the given ID
func PostOfferingsIDImageHandler(params operations.PostOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	//uploadImage("test", params.Upfile)

	return middleware.NotImplemented("")
}

// GetOfferingsIDImageHandler provides the Image for a Offering with a given ID
func GetOfferingsIDImageHandler(params operations.GetOfferingsIDImageParams, principal *models.Principal) middleware.Responder {

	/*file_ext := filepath.Ext(filename)

	content_type := mime.TypeByExtension(file_ext)

	if gridFSfile.Size() > 0 {
		r.Header.Set("Content-Type", content_type)
	} else {
		r.Header.Set("Content-Type", "image/*")
	}*/



	return middleware.NotImplemented("")
}
