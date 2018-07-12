package offerings

import (
	"testing"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
)

func TestGetOfferingByIDHandler(t *testing.T) {
	offeringName := "testName"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	okResponse := response.(*offerings.CreateNewOfferingOK)

	//log.Print(okResponse.Payload.ID.Hex())

	inOut := []struct {
		in  string
		out middleware.Responder
	}{
		{
			in:  "404040404040404040404040",
			out: &offerings.GetOfferingByIDNotFound{},
		},
		{
			in:  okResponse.Payload.ID.Hex(),
			out: &offerings.GetOfferingByIDOK{},
		},
	}

	for _, iot := range inOut {

		response := GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: iot.in}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}

	response = GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: okResponse.Payload.ID.Hex()}, testPrincipal)
	assert.IsType(t, &offerings.GetOfferingByIDOK{}, response)
	out := response.(*offerings.GetOfferingByIDOK)

	assert.Equal(t, out.Payload.ID, okResponse.Payload.ID)

}
