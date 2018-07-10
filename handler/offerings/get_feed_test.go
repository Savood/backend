package offerings

import (
	"testing"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
	"log"
)

func TestGetFeedHandler(t *testing.T) {
	offeringName := "testName"

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{
		Body: &models.Offering{
			Name:     offeringName,
			Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{47.4874592, 46.4660395}},
		}},
		testPrincipal,
	)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{
		Body: &models.Offering{
			Name:     offeringName,
			Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{47.4874591, 46.4660394}},
		}},
		testPrincipal,
	)
	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{
		Body: &models.Offering{
			Name:     offeringName,
			Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{47.4874592, 46.4660395}},
		}},
		&models.Principal{Userid: bson.ObjectIdHex("404040404040404040404040")},
	)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)
	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{
		Body: &models.Offering{
			Name:     offeringName,
			Location: &models.OfferingLocation{Type: "Point", Coordinates: []float64{47.4874591, 46.4660394}},
		}},
		&models.Principal{Userid: bson.ObjectIdHex("404040404040404040404040")},
	)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	inOut := []struct {
		inLon      float64
		inLat      float64
		inDistance float64
		out        middleware.Responder
		outLen     int
	}{
		{
			inLon:      47.4874592,
			inLat:      46.4660395,
			inDistance: 0,
			out:        &offerings.GetFeedOK{},
			outLen:     1,
		},
		{
			inLon:      47.4874592,
			inLat:      46.4660395,
			inDistance: 1,
			out:        &offerings.GetFeedOK{},
			outLen:     2,
		},
		{
			inLat: 200,
			out:   &offerings.GetFeedBadRequest{},
		},
		{
			inLon: 200,
			out:   &offerings.GetFeedBadRequest{},
		},
	}

	for _, iot := range inOut {

		response := GetFeedHandler(offerings.GetFeedParams{Lat: iot.inLat, Lon: iot.inLon, Distance: iot.inDistance}, testPrincipal)

		assert.IsType(t, iot.out, response)

		ok, tr := response.(*offerings.GetFeedOK)

		log.Print(iot)

		if tr {
			assert.Equal(t, iot.outLen, len(ok.Payload))
		}

	}

}
