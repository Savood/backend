package image

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
)

func PostOfferingIDImageHandler(params operations.PostOfferingIDImageParams, principal *models.Principal) middleware.Responder {

	//uploadImage("test", params.Upfile)

	return middleware.NotImplemented("")
}

/*

func GetOfferingIDImageHandler() middleware.Responder {

	file_ext := filepath.Ext(filename)

	content_type := mime.TypeByExtension(file_ext)

	if gridFSfile.Size() > 0 {
		r.Header.Set("Content-Type", content_type)
	} else {
		r.Header.Set("Content-Type", "image/*")
	}

	return middleware.NotImplemented("")
}
*/
