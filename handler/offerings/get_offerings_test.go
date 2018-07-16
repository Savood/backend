package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"log"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
)

func TestGetOfferingsHandler(t *testing.T) {
	var offeringIDs []string
	createOfferings := []struct {
		Principal *models.Principal
		Savood    bool
	}{
		{
			Principal: testPrincipal1,
		},
		{
			Principal: testPrincipal1,
		},
		{
			Principal: forbiddenPrincipal,
			Savood: true,
		},
		{
			Principal: forbiddenPrincipal,
			Savood: true,
		},
		{
			Principal: forbiddenPrincipal,
			Savood: true,
		},
	}

	for _, cO := range createOfferings {

		response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{
			Body: &models.Offering{
				Name:     "test",
				Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 1}},
			}},
			cO.Principal,
		)

		assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
		offeringIDs = append(offeringIDs, response.(*offerings.CreateNewOfferingOK).Payload.ID.Hex())

		if cO.Savood {
			response = PlaceSavoodHandler(operations.PlaceSavoodParams{OfferingID: response.(*offerings.CreateNewOfferingOK).Payload.ID.Hex()}, testPrincipal1)
			assert.IsType(t, &operations.PlaceSavoodOK{}, response)
		}
	}


	inOut := []struct {
		Filter   string
		Expected middleware.Responder
		Length   int
	}{
		{
			Filter:   "owned",
			Expected: &offerings.GetOfferingsOK{},
			Length:   2,
		},
		{
			Filter:   "requested",
			Expected: &offerings.GetOfferingsOK{},
			Length:   3,
		},
		{
			Filter:   "somethingUnapropriate",
			Expected: &offerings.GetOfferingsBadRequest{},
		},
		{
			Filter:   "",
			Expected: &offerings.GetOfferingsOK{},
			Length:   5,
		},
	}

	for _, iot := range inOut {

		var f *string

		if iot.Filter != "" {
			f = &iot.Filter
		}

		response := GetOfferingsHandler(offerings.GetOfferingsParams{Filter: f}, testPrincipal1)

		assert.IsType(t, iot.Expected, response)

		errResponse, ok := response.(*offerings.GetOfferingsInternalServerError)
		if ok {
			log.Print(*errResponse.Payload.Message)
		}

		ofs, ok := response.(*offerings.GetOfferingsOK)
		if ok {
			assert.Equal(t, iot.Length, len(ofs.Payload))
		}

	}
	for _, id := range offeringIDs {
		assert.NoError(t, dao.DeleteOfferingByID(id))
	}


}
