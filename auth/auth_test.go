package auth

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
)


const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtsYXVzQGNoZC5jeCIsImV4cCI6MjUyNDYwODAwMCwidXNlcmlkIjoiNWIxZmM4NmVmMDM1OTUwMDAxZDRhYThiIn0.NOE5QCjOvAJjCVio7_9wN62T7ywE6-ri3Mf7BMu-wvY"
const secret = "secret"


func TestGetAuthorizer(t *testing.T) {

	s := secret

	authFunc := GetAuthorizer(&s)

	principal, err := authFunc(token)

	assert.NoError(t, err)
	assert.NotNil(t, principal)

	if principal != nil {
		assert.IsType(t, models.Principal{}, *principal)

		assert.Equal(t, models.Principal{
			Email:    "klaus@chd.cx",
			Userid:   bson.ObjectIdHex("5b1fc86ef035950001d4aa8b"),
		}, *principal)
	}


}
