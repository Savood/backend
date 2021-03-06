package dao

import (
	"github.com/go-openapi/strfmt"
	"testing"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/database"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/stretchr/testify/assert"
	"time"
	"log"
	"fmt"
	"os"
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
			Zip:    "68167",
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	err := SaveUser(user)
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

	SaveOffering(offering)

	return offeringID, offering
}

func TestGetUserByID(t *testing.T) {
	userID, _ := CreateFakeUser()

	id, e := GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
}

func TestGetUserShortByID(t *testing.T) {
	userID, _ := CreateFakeUser()

	id, e := GetUserShortByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
}

func TestSaveUser(t *testing.T) {
	userID := bson.NewObjectId()

	user := &models.User{
		ID: userID,
		Address: &models.Address{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    "68167",
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	assert.NoError(t, SaveUser(user))

	id, e := GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
	assert.Equal(t, user, id)
}

func TestSaveSavood(t *testing.T) {
	savoodID, _ := CreateFakeOffering()
	userID := bson.NewObjectId()

	user := &models.User{
		ID: userID,
		Address: &models.Address{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    "68167",
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	assert.NoError(t, SaveUser(user))

	id, e := GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
	assert.Equal(t, user, id)

	err := AddSavoodToUserID(userID.Hex(), savoodID.Hex())
	assert.NoError(t, err)

	savoods, err := GetSavoodsByUserID(userID)
	assert.NoError(t, err)
	assert.Contains(t, savoods, savoodID)

	id, e = GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
	assert.Equal(t, user, id)

	assert.NoError(t, SaveUser(user))

	savoods, err = GetSavoodsByUserID(userID)
	assert.NoError(t, err)
	assert.Contains(t, savoods, savoodID)

}

func TestRemoveSavoodFromUserID(t *testing.T) {
	savoodID := bson.NewObjectId()
	userID := bson.NewObjectId()

	user := &models.User{
		ID: userID,
		Address: &models.Address{
			City:   "City",
			Number: "Number",
			Street: "Street",
			Zip:    "68167",
		},
		Badges:      []string{"badge1", "badge2"},
		Description: "description",
		Email:       "email@example.com",
		Firstname:   "Hans",
		Lastname:    "Peter",
		Phone:       "+49000000000",
	}

	assert.NoError(t, SaveUser(user))

	id, e := GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
	assert.Equal(t, user, id)

	err := AddSavoodToUserID(userID.Hex(), savoodID.Hex())
	assert.NoError(t, err)

	savoods, err := GetSavoodsByUserID(userID)
	assert.NoError(t, err)
	assert.Contains(t, savoods, savoodID)

	err = RemoveSavoodFromUserID(userID.Hex(), savoodID.Hex())
	assert.NoError(t, err)

	savoods, err = GetSavoodsByUserID(userID)
	assert.NoError(t, err)
	assert.NotContains(t, savoods, savoodID)

	id, e = GetUserByID(userID.Hex())
	assert.NotNil(t, id)
	assert.NoError(t, e)
	assert.Equal(t, user, id)

}

func TestGetAllChatsByUserID(t *testing.T) {
	userID, _ := CreateFakeUser()

	userShort, _ := GetUserShortByID(userID.Hex())

	offeringID, _ := CreateFakeOffering()

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID},
	}

	assert.NoError(t, SaveChat(&chat))

	chats, err := GetAllChatsByUserID(userID.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, chats)
}

func TestSaveChat(t *testing.T) {
	userID, _ := CreateFakeUser()

	userShort, _ := GetUserShortByID(userID.Hex())

	offeringID, _ := CreateFakeOffering()

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID},
	}

	assert.NoError(t, SaveChat(&chat))

	chatByID, err := GetChatByID(chatID.Hex(), userID.Hex())
	assert.NotNil(t, chatByID)
	assert.NoError(t, err)
}

func TestGetChatByID(t *testing.T) {
	TestSaveChat(t)
}

func TestGetAllMessagesByChatID(t *testing.T) {
	userID, _ := CreateFakeUser()

	userShort, _ := GetUserShortByID(userID.Hex())

	offeringID, _ := CreateFakeOffering()

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID},
	}

	assert.NoError(t, SaveChat(&chat))

	SaveMessage(chat, models.Message{
		From:    userShort,
		Content: "hello",
		Time:    strfmt.DateTime(time.Now().UTC()),
	})

	messages, err := GetAllMessagesByChatID(chatID.Hex())
	assert.NoError(t, err)
	assert.True(t, len(messages) > 0)

	if len(messages) > 0 {
		assert.Equal(t, "hello", messages[0].Content)
	}
}

func TestSaveMessage(t *testing.T) {
	TestGetAllMessagesByChatID(t)
}

func TestGetAllOfferingsByUserID(t *testing.T) {
	_, offering := CreateFakeOffering()

	offerings, err := GetAllOfferingsByUserID(offering.Creator.ID.Hex(), nil)
	assert.NoError(t, err)
	assert.True(t, len(offerings) > 0)

	if len(offerings) > 0 {
		assert.Equal(t, "description", offerings[0].Description)
	}
}

func TestGetOfferingByID(t *testing.T) {
	offeringID, _ := CreateFakeOffering()

	offering, err := GetOfferingByID(offeringID.Hex(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, offering)

	assert.Equal(t, "description", offering.Description)
}

func TestGetNearOfferings(t *testing.T) {
	_, offering := CreateFakeOffering()

	offerings, err := GetNearOfferings(*offering.Location, 200, nil)

	assert.NoError(t, err)
	assert.NotNil(t, offerings)
	assert.True(t, len(offerings) > 0)
}

func TestSaveOffering(t *testing.T) {
	offeringID, _ := CreateFakeOffering()

	offering, err := GetOfferingByID(offeringID.Hex(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, offering)

	assert.Equal(t, "description", offering.Description)
}

func TestGetAllChatsByOfferingAndUserID(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	userID, _ := CreateFakeUser()

	userShort, _ := GetUserShortByID(userID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	SaveChat(&chat)

	chatsPartner, err := GetAllChatsByOfferingAndUserID(offeringTest.ID.Hex(), userID.Hex())
	assert.NoError(t, err)
	assert.True(t, len(chatsPartner) > 0)

	chatsCreator, err := GetAllChatsByOfferingAndUserID(offeringTest.ID.Hex(), offeringTest.Creator.ID.Hex())
	assert.NoError(t, err)
	assert.True(t, len(chatsCreator) > 0)
}

func TestGetChatByIDAndUserID(t *testing.T) {
	userID, _ := CreateFakeUser()

	offeringID, _ := CreateFakeOffering()

	userShort, _ := GetUserShortByID(userID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringID},
		Partner:    userShort,
	}

	SaveChat(&chat)

	chatl, err := GetChatByIDAndUserID(chat.ID.Hex(), userID.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, chatl)

	chatl, err = GetChatByIDAndUserID(chat.ID.Hex(), bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	assert.Nil(t, chatl)
}

func TestUpdateChatRemoveOfferingID(t *testing.T) {
	userID, _ := CreateFakeUser()

	offeringID, _ := CreateFakeOffering()

	userShort, _ := GetUserShortByID(userID.Hex())

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID},
	}

	assert.NoError(t, SaveChat(&chat))

	UpdateChatRemoveOfferingID(offeringID.Hex())

	chatByID, err := GetChatByID(chatID.Hex(), userID.Hex())
	assert.Nil(t, chatByID)
	assert.Error(t, err)
}

func TestUpdateChatRemoveOfferingID2(t *testing.T) {
	userID, _ := CreateFakeUser()

	offeringID, _ := CreateFakeOffering()

	userShort, _ := GetUserShortByID(userID.Hex())

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID, bson.NewObjectId()},
	}

	assert.NoError(t, SaveChat(&chat))

	UpdateChatRemoveOfferingID(offeringID.Hex())

	chatByID, err := GetChatByID(chatID.Hex(), userID.Hex())
	assert.NotNil(t, chatByID)
	assert.NoError(t, err)
}

func TestTouchChat(t *testing.T) {
	userID, _ := CreateFakeUser()

	offeringID, _ := CreateFakeOffering()

	userShort, _ := GetUserShortByID(userID.Hex())

	chatID := bson.NewObjectId()

	chat := models.Chat{
		ID:         chatID,
		Partner:    userShort,
		OfferingID: []bson.ObjectId{offeringID, bson.NewObjectId()},
	}

	assert.NoError(t, SaveChat(&chat))

	assert.NoError(t, TouchChat(chat.ID.Hex()))
}
