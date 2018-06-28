package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
	"time"
	"git.dhbw.chd.cx/savood/backend/dao"
)

// CreateNewOfferingHandler creates an offering and ensures it is in order
func CreateNewOfferingHandler(params offerings.CreateNewOfferingParams, principal *models.Principal) middleware.Responder {

	if params.Body == nil {
		return offerings.NewCreateNewOfferingBadRequest()
	}

	offering := *params.Body

	if offering.Name == "" ||
		offering.Location == nil || offering.Location.Coordinates == nil ||
		offering.Location.Type == "" || len(offering.Location.Coordinates) != 2 {

		return offerings.NewCreateNewOfferingBadRequest()

	}

	if offering.ID != bson.ObjectId("") ||
		!(offering.CreatorID == bson.ObjectId("") || offering.CreatorID == principal.Userid) ||
		offering.Time != strfmt.DateTime(time.Time{}) || offering.RequestedBy != 0 {

		return offerings.NewCreateNewOfferingForbidden()
	}

	offering.CreatorID = principal.Userid
	offering.Time = strfmt.DateTime(time.Now())
	offering.RequestedBy = 0

	dao.SaveOffering(&offering)

	return offerings.NewCreateNewOfferingOK().WithPayload(&offering)
}
