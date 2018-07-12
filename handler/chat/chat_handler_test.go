package chat

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"os"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/messages"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo/bson"
	"fmt"
	"github.com/go-openapi/strfmt"
	"time"
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
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

func TestMessagesCreateNewMessageHandler(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	_, sender := CreateFakeUser()

	userShort, _ := dao.GetUserShortByID(sender.ID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	dao.SaveChat(&chat)

	principal := models.Principal{
		Email:  sender.Email.String(),
		Userid: sender.ID,
	}

	params := messages.CreateNewMessageParams{
		ChatID: chat.ID.Hex(),
		Body: &models.Message{
			Content: "Hello?!?",
		},
	}

	response := MessagesCreateNewMessageHandler(params, &principal)
	assert.IsType(t, &messages.CreateNewMessageOK{}, response)
}

func TestMessagesCreateNewMessageHandler2(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	_, sender := CreateFakeUser()

	userShort, _ := dao.GetUserShortByID(sender.ID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	dao.SaveChat(&chat)

	principal := models.Principal{
		Email:  sender.Email.String(),
		Userid: bson.NewObjectId(),
	}

	params := messages.CreateNewMessageParams{
		ChatID: chat.ID.Hex(),
		Body: &models.Message{
			Content: "Hello?!?",
		},
	}

	response := MessagesCreateNewMessageHandler(params, &principal)
	assert.IsType(t, &messages.CreateNewMessageBadRequest{}, response)
}

func TestMessagesGetAllChatsForOfferingHandler(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	_, sender := CreateFakeUser()

	userShort, _ := dao.GetUserShortByID(sender.ID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	dao.SaveChat(&chat)

	params := operations.GetAllChatsForOfferingParams{
		ID: offeringTest.ID.Hex(),
	}

	principal := models.Principal{
		Email:  sender.Email.String(),
		Userid: userShort.ID,
	}

	response := MessagesGetAllChatsForOfferingHandler(params, &principal)
	assert.IsType(t, &operations.GetAllChatsForOfferingOK{}, response)
}

func TestMessagesGetAllMessagesForChatHandler(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	_, sender := CreateFakeUser()

	userShort, _ := dao.GetUserShortByID(sender.ID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	dao.SaveChat(&chat)

	message := models.Message{
		Content: "Hello?!?",
		From:    userShort,
		Time:    strfmt.NewDateTime(),
	}

	dao.SaveMessage(chat, message)

	principal := models.Principal{
		Email:  sender.Email.String(),
		Userid: userShort.ID,
	}

	params := messages.GetAllMessagesForChatParams{
		ChatID: chat.ID.Hex(),
	}

	response := MessagesGetAllMessagesForChatHandler(params, &principal)
	assert.IsType(t, &messages.GetAllMessagesForChatOK{}, response)
}

func TestMessagesGetAllChatsHandler(t *testing.T) {
	_, offeringTest := CreateFakeOffering()

	_, sender := CreateFakeUser()

	userShort, _ := dao.GetUserShortByID(sender.ID.Hex())

	chat := models.Chat{
		OfferingID: []bson.ObjectId{offeringTest.ID},
		Partner:    userShort,
	}

	dao.SaveChat(&chat)

	principal := models.Principal{
		Email:  sender.Email.String(),
		Userid: userShort.ID,
	}

	response := MessagesGetAllChatsHandler(messages.GetAllChatsParams{}, &principal)
	assert.IsType(t, &messages.GetAllChatsOK{}, response)
}
