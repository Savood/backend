package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/runtime/middleware"
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
	"time"
)

func TestUpdateOfferingByIDHandler(t *testing.T) {
	offeringName := "testName"
	changedText := "changedText"
	location := &models.OfferingLocation{Type: "Point", Coordinates: []float64{1, 2}}

	response := CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, testPrincipal)
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	ok := response.(*offerings.CreateNewOfferingOK)

	response = CreateNewOfferingHandler(offerings.CreateNewOfferingParams{Body: &models.Offering{Name: offeringName, Location: location}}, &models.Principal{Userid: bson.ObjectIdHex("404040404040404040404040")})
	assert.IsType(t, &offerings.CreateNewOfferingOK{}, response)

	forbidden := response.(*offerings.CreateNewOfferingOK)

	inOut := []struct {
		in   *models.Offering
		inID bson.ObjectId
		out  middleware.Responder
	}{
		{
			inID: ok.Payload.ID,
			out:  &offerings.UpdateOfferingByIDNoContent{},
		},
		{
			inID: forbidden.Payload.ID,
			out:  &offerings.UpdateOfferingByIDForbidden{},
		},
		{
			inID: bson.ObjectIdHex("404040404040404040404040"),
			out:  &offerings.UpdateOfferingByIDNotFound{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{ID: bson.ObjectId("somethingnotnothing")},
			out:  &offerings.UpdateOfferingByIDForbidden{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{ID: ok.Payload.ID},
			out:  &offerings.UpdateOfferingByIDNoContent{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{Creator: &models.UserShort{ID: bson.ObjectId("somethingnotnothing")}},
			out:  &offerings.UpdateOfferingByIDForbidden{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{Time: strfmt.DateTime(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC))},
			out:  &offerings.UpdateOfferingByIDForbidden{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{Location: &models.OfferingLocation{}},
			out:  &offerings.UpdateOfferingByIDNoContent{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{Location: &models.OfferingLocation{Coordinates: []float64{1, 2, 3}}},
			out:  &offerings.UpdateOfferingByIDBadRequest{},
		},
		{
			inID: ok.Payload.ID,
			in:   &models.Offering{RequestedBy: 99, Description: changedText},
			out:  &offerings.UpdateOfferingByIDNoContent{},
		},
	}

	for _, iot := range inOut {

		response := UpdateOfferingByIDHandler(offerings.UpdateOfferingByIDParams{ID: iot.inID.Hex(), Body: iot.in}, testPrincipal)

		assert.IsType(t, iot.out, response)

	}

	response = GetOfferingByIDHandler(offerings.GetOfferingByIDParams{ID: ok.Payload.ID.Hex()}, testPrincipal)
	assert.IsType(t, &offerings.GetOfferingByIDOK{}, response)
	out := response.(*offerings.GetOfferingByIDOK)

	assert.Equal(t, out.Payload.ID, ok.Payload.ID)
	assert.Equal(t, changedText, out.Payload.Description)
	assert.Equal(t, location.Type, out.Payload.Location.Type)
	assert.Equal(t, int64(0), out.Payload.RequestedBy)

}
