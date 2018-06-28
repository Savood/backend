package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo"
	"github.com/go-openapi/strfmt"
	"time"
)

func UpdateOfferingByIDHandler(params offerings.UpdateOfferingByIDParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.ID)

	if err == mgo.ErrNotFound {
		return offerings.NewUpdateOfferingByIDNotFound()
	}

	if err != nil || offering == nil {
		str := err.Error()
		return offerings.NewUpdateOfferingByIDInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.CreatorID != principal.Userid {
		return offerings.NewUpdateOfferingByIDForbidden()
	}

	var emptyOffering models.Offering

	if params.Body == nil || *params.Body == emptyOffering {
		return offerings.NewUpdateOfferingByIDNoContent()
	}

	patch := *params.Body

	if !(patch.CreatorID == principal.Userid || patch.CreatorID == "") ||
		!(patch.ID == offering.ID || patch.ID == "") ||
		!(patch.Time == offering.Time || patch.Time == strfmt.DateTime(time.Time{})) {

		return offerings.NewUpdateOfferingByIDForbidden()

	}

	if patch.Location != nil && !(len(patch.Location.Coordinates) == 2 || len(patch.Location.Coordinates) == 0) {
		return offerings.NewUpdateOfferingByIDBadRequest()
	}

	patchOffering(offering, patch)

	err = dao.SaveOffering(offering)

	if err != nil {
		str := err.Error()
		return offerings.NewUpdateOfferingByIDInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return offerings.NewUpdateOfferingByIDNoContent()
}

func patchOffering(offering *models.Offering, patch models.Offering) {
	if patch.Location != nil {
		if len(patch.Location.Coordinates) > 0 {
			offering.Location.Coordinates = patch.Location.Coordinates
		}
		if len(patch.Location.Type) > 0 {
			offering.Location.Type = patch.Location.Type
		}
	}

	if len(patch.Description) > 0 {
		offering.Description = patch.Description
	}

	if len(patch.Name) > 0 {
		offering.Name = patch.Name
	}

	if patch.Address != nil {
		if len(patch.Address.City) > 0 {
			offering.Address.City = patch.Address.City

		}
		if len(patch.Address.Number) > 0 {
			offering.Address.Number = patch.Address.Number

		}
		if len(patch.Address.Street) > 0 {
			offering.Address.Street = patch.Address.Street

		}
		if patch.Address.Zip != 0 {
			offering.Address.Zip = patch.Address.Zip
		}
	}

	if patch.BestByDate != strfmt.Date(time.Time{}) {

		offering.BestByDate = patch.BestByDate
	}
}
