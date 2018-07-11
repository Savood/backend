package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
)


// GetOfferingsHandler handles getting all offerings, which are owned or savooded
func GetOfferingsHandler(params offerings.GetOfferingsParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation offerings.GetOfferings has not yet been implemented")
}
