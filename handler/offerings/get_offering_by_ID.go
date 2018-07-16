package offerings

import (
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo"
)

// GetOfferingByIDHandler gets an offering by id
func GetOfferingByIDHandler(params offerings.GetOfferingByIDParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.ID, principal)

	if err == mgo.ErrNotFound {
		return offerings.NewGetOfferingByIDNotFound()
	}

	if err != nil || offering == nil {
		str := err.Error()
		return offerings.NewGetOfferingByIDInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return offerings.NewGetOfferingByIDOK().WithPayload(offering)

}
