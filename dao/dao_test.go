package dao

import (
	"github.com/go-openapi/strfmt"

	"testing"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/database"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
)

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		fmt.Println("Failed DB connection")
		return
	}
	m.Run()
}

func CreateFakeUser() (bson.ObjectId, *models.User) {
	userID := bson.NewObjectId()

	user := &models.User{
		ID: string(userID),
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

	SaveUser(user)

	return userID, user
}

func CreateFakeOffering() (bson.ObjectId, *models.Offering) {
	userID, _ := CreateFakeUser()

	offeringID := bson.NewObjectId()

	offering := &models.Offering{
		Time:        strfmt.DateTime(time.Now().UTC()),
		CreatorID:   string(userID),
		ID:          string(offeringID),
		Description: "description",
		Name:        "name",
		AvatarURL:   "avatar-url",
		BestByDate:  strfmt.Date(time.Now().UTC()),
		Location: &models.OfferingLocation{
			Coordinates: []float64{0.0, 0.0},
			Type:        "type",
		},
		RequestedBy: 0,
	}

	SaveOffering(offering)

	return offeringID, offering
}

func TestGetUserByID(t *testing.T) {
	userID, _ := CreateFakeUser()

	id, e := GetUserByID(string(userID))
	assert.NotNil(t, id)
	assert.NoError(t, e)
}

func TestGetUserShortByID(t *testing.T) {
	userID, _ := CreateFakeUser()

	id, e := GetUserShortByID(string(userID))
	assert.NotNil(t, id)
	assert.NoError(t, e)
}

func TestSaveUser(t *testing.T) {
	userID := bson.NewObjectId()

	user := &models.User{
		ID: string(userID),
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

	id, e := GetUserByID(string(userID))
	assert.NotNil(t, id)
	assert.NoError(t, e)
}

func TestGetAllChatsByUserID(t *testing.T) {
	userID, _ := CreateFakeUser()

	chats, err := GetAllChatsByUserID(string(userID))
	assert.NoError(t, err)
	assert.NotNil(t, chats)
}

func TestSaveChat(t *testing.T) {
	userID, user := CreateFakeUser()

	userShort, _ := GetUserShortByID(string(userID))

	principal := models.Principal{
		Email:    string(user.Email),
		Userid:   string(userID),
		Username: "",
	}

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         string(chatID),
		Partner:    userShort,
		OfferingID: []string{},
	}

	assert.NoError(t, SaveChat(principal, chat))

	chatByID, err := GetChatByID(string(chatID))
	assert.NotNil(t, chatByID)
	assert.NoError(t, err)
}

func TestGetChatByID(t *testing.T) {
	TestSaveChat(t)
}

func TestGetAllMessagesByChatID(t *testing.T) {
	userID, user := CreateFakeUser()

	userShort, _ := GetUserShortByID(string(userID))

	principal := models.Principal{
		Email:    string(user.Email),
		Userid:   string(userID),
		Username: "",
	}

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         string(chatID),
		Partner:    userShort,
		OfferingID: []string{},
	}

	assert.NoError(t, SaveChat(principal, chat))

	SaveMessage(chat, models.Message{
		From:    userShort,
		Content: "hello",
		Time:    strfmt.DateTime(time.Now().UTC()),
	})

	messages, err := GetAllMessagesByChatID(string(chatID))
	assert.NoError(t, err)
	assert.True(t, len(messages) > 1)

	assert.Equal(t, "hello", messages[0].Content)
}

func TestSaveMessage(t *testing.T) {
	TestGetAllMessagesByChatID(t)
}

func TestGetAllOfferingsByUserID(t *testing.T) {
	_, offering := CreateFakeOffering()

	offerings, err := GetAllOfferingsByUserID(offering.CreatorID)
	assert.NoError(t, err)
	assert.True(t, len(offerings) > 1)

	assert.Equal(t, "description", offerings[0].Description)
}

func TestGetOfferingByID(t *testing.T) {
	offeringID, _ := CreateFakeOffering()

	offering, err := GetOfferingByID(string(offeringID))
	assert.NoError(t, err)
	assert.NotNil(t, offering)

	assert.Equal(t, "description", offering.Description)
}

func TestSaveOffering(t *testing.T) {
	offeringID, _ := CreateFakeOffering()

	offering, err := GetOfferingByID(string(offeringID))
	assert.NoError(t, err)
	assert.NotNil(t, offering)

	assert.Equal(t, "description", offering.Description)
}
