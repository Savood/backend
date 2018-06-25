package auth

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
)


const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtsYXVzQGNoZC5jeCIsImV4cCI6MTUyOTkxNjI1MywidXNlcmlkIjoiNWIxZmM4NmVmMDM1OTUwMDAxZDRhYThiIn0.aZY84deyYN_-JsSlPdk0GyArlwonwLPBQTQMhkKMJSI"

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
			Userid:   "5b1fc86ef035950001d4aa8b",
		}, *principal)
	}


}
