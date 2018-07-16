package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/runtime/middleware"
)

func TestRemoveSavoodHandler(t *testing.T) {
	offeringName := "testName"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	forbiddenOffering := response.(*offerings.CreateNewOfferingOK)

	//log.Print(okOffering.Payload.ID)

	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, forbiddenPrincipal)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	notForbiddenOffering := response.(*offerings.CreateNewOfferingOK)

	response = PlaceSavoodHandler(operations.PlaceSavoodParams{OfferingID: notForbiddenOffering.Payload.ID.Hex()}, testPrincipal)
	assert.IsType(t, &operations.PlaceSavoodOK{}, response)

	response = GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: notForbiddenOffering.Payload.ID.Hex()}, forbiddenPrincipal)
	assert.IsType(t, &offerings.GetOfferingByIDOK{}, response)
	okOffering := response.(*offerings.GetOfferingByIDOK)

	assert.Equal(t, int64(1), okOffering.Payload.RequestedBy)
	//log.Print(forbiddenOffering.Payload.ID)

	inOut := []struct {
		offeringID string
		expected   middleware.Responder
	}{
		{
			offeringID: forbiddenOffering.Payload.ID.Hex(),
			expected:   &operations.UnSavoodForbidden{},
		},
		{
			offeringID: notForbiddenOffering.Payload.ID.Hex(),
			expected:   &operations.UnSavoodOK{},
		},
		{
			offeringID: notForbiddenOffering.Payload.ID.Hex(),
			expected:   &operations.UnSavoodOK{},
		},
		{
			offeringID: "404040404040404040404040",
			expected:   &operations.UnSavoodNotFound{},
		},
	}

	for _, iot := range inOut {

		response := UnSavoodHandler(operations.UnSavoodParams{OfferingID: iot.offeringID}, testPrincipal)

		assert.IsType(t, iot.expected, response)

	}

	response = GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: notForbiddenOffering.Payload.ID.Hex()}, forbiddenPrincipal)
	assert.IsType(t, &offerings.GetOfferingByIDOK{}, response)
	okOffering = response.(*offerings.GetOfferingByIDOK)

	assert.Equal(t, int64(0), okOffering.Payload.RequestedBy)

}
