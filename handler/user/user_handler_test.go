package user

import (
	"git.dhbw.chd.cx/savood/backend/dao"
	"fmt"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/strfmt"
	"time"
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"os"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}
	os.Exit(m.Run())
}

func CreateFakeUser() (bson.ObjectId, *models.User) {
	userID := bson.NewObjectId()

	user := &models.User{
		ID: userID,
		Address: &models.Address{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    68167,
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	err := dao.SaveUser(user)
	if err != nil {
		fmt.Println("Failed fake user", err)
	}

	return userID, user
}

func CreateFakeOffering() (bson.ObjectId, *models.Offering) {
	userID, _ := CreateFakeUser()

	offeringID := bson.NewObjectId()

	offering := &models.Offering{
		Time: strfmt.DateTime(time.Now().UTC()),
		Creator: &models.UserShort{
			ID: userID,
		},
		ID:          offeringID,
		Description: "description",
		Name:        "name",
		BestByDate:  strfmt.Date(time.Now().UTC()),
		Location: &models.OfferingLocation{
			Coordinates: []float64{0.0, 0.0},
			Type:        "Point",
		},
		RequestedBy: 0,
	}

	dao.SaveOffering(offering)

	return offeringID, offering
}

func TestUsersCreateNewUserHandler(t *testing.T) {
	userID := bson.NewObjectId()

	user := &models.User{
		ID: userID,
		Address: &models.Address{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    68167,
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	params := users.CreateNewUserParams{
		Body: user,
	}

	principal := models.Principal{
		Userid: userID,
		Email:  user.Email.String(),
	}

	response := UsersCreateNewUserHandler(params, &principal)
	assert.IsType(t, &users.CreateNewUserOK{}, response)

	user, err := dao.GetUserByID(user.ID.Hex())
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUsersDeleteUserByIDHandler(t *testing.T) {
	userID, user := CreateFakeUser()

	params := users.DeleteUserByIDParams{
		ID: userID.Hex(),
	}

	principal := models.Principal{
		Userid: userID,
		Email:  user.Email.String(),
	}

	response := UsersDeleteUserByIDHandler(params, &principal)
	assert.IsType(t, &users.DeleteUserByIDNoContent{}, response)
}

func TestUsersGetUserByIDHandler(t *testing.T) {
	userID, user := CreateFakeUser()

	params := users.GetUserByIDParams{
		ID: userID.Hex(),
	}

	principal := models.Principal{
		Userid: userID,
		Email:  user.Email.String(),
	}

	response := UsersGetUserByIDHandler(params, &principal)
	assert.IsType(t, &users.GetUserByIDOK{}, response)
}

func TestUsersUpdateUserByIDHandler(t *testing.T) {
	userID, user := CreateFakeUser()

	user.Firstname = "Günther"

	params := users.UpdateUserByIDParams{
		Body: user,
	}

	principal := models.Principal{
		Userid: userID,
		Email:  user.Email.String(),
	}

	response := UsersUpdateUserByIDHandler(params, &principal)
	assert.IsType(t, &users.UpdateUserByIDNoContent{}, response)
}

func TestUsersUpdateUserByIDHandler2(t *testing.T) {
	_, user := CreateFakeUser()

	user.Firstname = "Günther"

	params := users.UpdateUserByIDParams{
		Body: user,
	}

	principal := models.Principal{
		Userid: bson.NewObjectId(),
		Email:  user.Email.String(),
	}

	response := UsersUpdateUserByIDHandler(params, &principal)
	assert.IsType(t, &users.UpdateUserByIDBadRequest{}, response)
}
