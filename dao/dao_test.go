package dao

import (
	"testing"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/database"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	database.ConnectDatabase(nil, nil)
}

func TestGetUserByID(t *testing.T) {
	userId := bson.NewObjectId()

	user := &models.User{
		ID: string(userId),
		Address: &models.UserAddress{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    68167,
		},
		AvatarID:     "avatarid",
		BackgroundID: "backgroundid",
		Badges:       []string{"badge1", "badge2"},
		Description:  "description",
		Email:        "email@example.com",
		Firstname:    "Hans",
		Lastname:     "Peter",
		Phone:        "+49000000000",
	}

	assert.NoError(t, SaveUser(user))

	id, e := GetUserByID(string(userId))
	assert.NotNil(t, id)
	assert.NoError(t, e)
}
