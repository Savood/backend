package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
)

// GetOfferingsHandler handles getting all offerings, which are owned or savooded
func GetOfferingsHandler(params offerings.GetOfferingsParams, principal *models.Principal) middleware.Responder {

	owned := true
	requested := true

	if params.Filter != nil {
		if *params.Filter == "owned" {
			requested = false
		} else if *params.Filter == "requested" {
			owned = false
		} else {
			return offerings.NewGetOfferingsBadRequest()
		}
	}

	var payload []*models.Offering

	if owned {
		o, err := dao.GetAllOfferingsByUserID(principal.Userid.Hex(), principal)
		if err != nil {
			str := err.Error()
			return offerings.NewGetOfferingsInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
		payload = append(payload, o...)
	}

	if requested {
		ids , err := dao.GetSavoodsByUserID(principal.Userid)
		if err != nil {
			str := err.Error()
			return offerings.NewGetOfferingsInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
		o, err := dao.GetOfferingsByIDs(ids, principal)
		if err != nil {
			str := err.Error()
			return offerings.NewGetOfferingsInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
		payload = append(payload, o...)
	}

	return offerings.NewGetOfferingsOK().WithPayload(payload)
}
