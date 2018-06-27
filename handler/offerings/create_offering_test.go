package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/strfmt"
	"time"
)

func TestOfferingsCreateNewOfferingHandler(t *testing.T) {

	offeringName := "testName"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	inOut := []struct {
		in  *models.Offering
		out middleware.Responder
	}{
		{
			in:  &models.Offering{},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Name: offeringName,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Location: location,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Location: &models.OfferingLocation{},
				Name:     offeringName,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Location: &models.OfferingLocation{Type: "Point"},
				Name:     offeringName,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Location: &models.OfferingLocation{Coordinates: []float64{1, 2}},
				Name:     offeringName,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},
		{
			in: &models.Offering{
				Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2, 3}},
				Name:     offeringName,
			},
			out: &offerings.CreateNewOfferingBadRequest{},
		},

		{
			in: &models.Offering{
				ID:       bson.ObjectId("somethingnotnothing"),
				Name:     offeringName,
				Location: location,
			},
			out: &offerings.CreateNewOfferingForbidden{},
		},
		{
			in: &models.Offering{
				CreatorID: bson.ObjectId("somethingotherthaninprincipal"),
				Name:      offeringName,
				Location:  location,
			},
			out: &offerings.CreateNewOfferingForbidden{},
		},
		{
			in: &models.Offering{
				RequestedBy: 99,
				Name:        offeringName,
				Location:    location,
			},
			out: &offerings.CreateNewOfferingForbidden{},
		},
		{
			in: &models.Offering{
				Time:     strfmt.DateTime(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)),
				Name:     offeringName,
				Location: location,
			},
			out: &offerings.CreateNewOfferingForbidden{},
		},
	}

	for _, iot := range inOut {

		response := OfferingsCreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: iot.in}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}

	response := OfferingsCreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)

	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	ok := response.(*offerings.CreateNewOfferingOK)

	assert.IsType(t, &models.Offering{}, ok.Payload)

	assert.IsType(t, bson.ObjectId(""), ok.Payload.ID)

	assert.Equal(t, testPrincipal.Userid, ok.Payload.CreatorID)

	assert.Equal(t, offeringName, ok.Payload.Name)

}
