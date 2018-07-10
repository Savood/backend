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
	okOffering := response.(*offerings.CreateNewOfferingOK)

	//log.Print(okOffering.Payload.ID)

	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, forbiddenPrincipal)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	forbiddenOffering := response.(*offerings.CreateNewOfferingOK)

	//log.Print(forbiddenOffering.Payload.ID)

	inOut := []struct {
		in  bson.ObjectId
		out middleware.Responder
	}{
		{
			in:  bson.ObjectIdHex("404040404040404040404040"),
			out: &offerings.DeleteOfferingByIDNoContent{},
		},
		{
			in:  okOffering.Payload.ID,
			out: &offerings.DeleteOfferingByIDNoContent{},
		},
		{
			in:  forbiddenOffering.Payload.ID,
			out: &offerings.DeleteOfferingByIDForbidden{},
		},
	}

	for _, iot := range inOut {

		response := DeleteOfferingByIDHandler(offerings.DeleteOfferingByIDParams{ID: iot.in.Hex()}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}
}
