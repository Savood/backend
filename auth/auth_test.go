package auth

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/errors"
)

const secret = "secret"

func TestGetAuthorizer(t *testing.T) {

	s := secret

	authFunc := GetAuthorizer(&s)

	InOut := []struct {
		Token             string
		ExpectedError     error
		ExpectedPrincipal *models.Principal
	}{
		{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtsYXVzQGNoZC5jeCIsImV4cCI6MjUyNDYwODAwMCwidXNlcmlkIjoiNWIxZmM4NmVmMDM1OTUwMDAxZDRhYThiIn0.NOE5QCjOvAJjCVio7_9wN62T7ywE6-ri3Mf7BMu-wvY",
			ExpectedPrincipal: &models.Principal{
				Email:  "klaus@chd.cx",
				Userid: bson.ObjectIdHex("5b1fc86ef035950001d4aa8b"),
			},
			ExpectedError: nil,
		},
		{
			Token:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtsYXVzQGNoZC5jeCIsImV4cCI6MjUyNDYwODAwMCwidXNlcmlkIjoiNWIxZmM4NmVmMDM1OTUwMDAxZDRhYThiIn0.eOtKhpaHfuO2OMLayYcT2nBQ50dcV94Sxz-4cgQtlcU",
			ExpectedPrincipal: nil,
			ExpectedError:     errors.New(401, "signature is invalid"),
		},
		{
			Token:             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtsYXVzQGNoZC5jeCIsImV4cCI6MSwidXNlcmlkIjoiNWIxZmM4NmVmMDM1OTUwMDAxZDRhYThiIn0.MA3Lom739v84hfjhPVh7Ai0mw2OIOjFJipdng7rL4kE",
			ExpectedPrincipal: nil,
			ExpectedError:     errors.New(401, "Token is expired"),
		},
	}

	for _, iot := range InOut {


		principal, err := authFunc(iot.Token)

		if iot.ExpectedPrincipal != nil {

			assert.Equal(t, *iot.ExpectedPrincipal, *principal)

		} else {
			assert.Nil(t, principal)
		}

		if iot.ExpectedError != nil {

			assert.Equal(t, iot.ExpectedError, err)

		} else {
			assert.Nil(t, err)
		}

	}

}
