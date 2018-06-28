package offerings

import (
	"testing"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
)

func TestGetOfferingByIDHandler(t *testing.T) {
	offeringName := "testName"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)

	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	ok := response.(*offerings.CreateNewOfferingOK)

	inOut := []struct {
		in  bson.ObjectId
		out middleware.Responder
	}{
		{
			in:  bson.ObjectIdHex("404040404040404040404040"),
			out: &offerings.GetOfferingByIDNotFound{},
		},
		{
			in:  ok.Payload.ID,
			out: &offerings.GetOfferingByIDOK{},
		},
	}

	for _, iot := range inOut {

		response := GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: iot.in.Hex()}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}

	response = GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: ok.Payload.ID.Hex()}, testPrincipal)

	assert.IsType(t, &offerings.GetOfferingByIDOK{}, response)

	out := response.(*offerings.GetOfferingByIDOK)

	assert.Equal(t, out.Payload.ID, ok.Payload.ID)

}
