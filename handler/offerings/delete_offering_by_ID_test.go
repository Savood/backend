package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/runtime/middleware"
	"github.com/globalsign/mgo/bson"
)

func TestDeleteOfferingByIDHandler(t *testing.T) {
	offeringName := "testName"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)

	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	ok := response.(*offerings.CreateNewOfferingOK)

	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, &models.Principal{Userid:bson.ObjectIdHex("404040404040404040404040")})

	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	forbidden := response.(*offerings.CreateNewOfferingOK)


	inOut := []struct {
		in  bson.ObjectId
		out middleware.Responder
	}{
		{
			in:  bson.ObjectIdHex("404040404040404040404040"),
			out: &offerings.DeleteOfferingByIDNoContent{},
		},
		{
			in:  ok.Payload.ID,
			out: &offerings.DeleteOfferingByIDNoContent{},
		},
		{
			in:  forbidden.Payload.ID,
			out: &offerings.DeleteOfferingByIDForbidden{},
		},
	}

	for _, iot := range inOut {

		response := DeleteOfferingByIDHandler(offerings.DeleteOfferingByIDParams{ID: iot.in.Hex()}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}
}
