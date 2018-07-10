package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo"
)

// DeleteOfferingByIDHandler handles deletion of offering by id
func DeleteOfferingByIDHandler(params offerings.DeleteOfferingByIDParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.ID)

	if err == mgo.ErrNotFound {
		return offerings.NewDeleteOfferingByIDNoContent()
	}

	if err != nil {
		str := err.Error()
		return offerings.NewDeleteOfferingByIDInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.Creator.ID != principal.Userid {
		return offerings.NewDeleteOfferingByIDForbidden()
	}

	err = dao.DeleteOfferingByID(params.ID)
	if err != nil {
		str := err.Error()
		return offerings.NewDeleteOfferingByIDInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}


	return offerings.NewDeleteOfferingByIDNoContent()
}
