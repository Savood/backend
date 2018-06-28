package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
)

func GetFeedHandler(params offerings.GetFeedParams, principal *models.Principal) middleware.Responder {

	if params.Lat > 90 || params.Lat < -90 ||
		params.Lon > 180 || params.Lon < -180 {
		return offerings.NewGetFeedBadRequest()
	}

	out, err := dao.GetNearOfferings(models.OfferingLocation{Type: "Point", Coordinates: []float64{params.Lon, params.Lat}}, params.Distance)

	if err != nil {
		str := err.Error()
		return offerings.NewGetFeedInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	b := out[:0]
	for _, x := range out {
		if x.CreatorID != principal.Userid {
			b = append(b, x)
		}
	}

	return offerings.NewGetFeedOK().WithPayload(b)

}
