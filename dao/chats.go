package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//ChatsCollectionName collection of chats in mongodb
const ChatsCollectionName = "chats"

//ChatTO Transfer Object for Chat
type ChatTO struct {
	ID string `json:"_id"`

	OfferingID []string `json:"offering-id"`

	Partner string `json:"partner"`

	OfferingCreatorID string `json:"offering-creator-id"`
}

//GetAllChatsByUserID Get All chats by user id (by offerings and partner)
func GetAllChatsByUserID(userID string) ([]*models.Chat, error) {
	offerings, err := GetAllOfferingsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var offeringIds []bson.ObjectId

	for _, result := range offerings {
		offeringIds = append(offeringIds, bson.ObjectIdHex(result.ID))
	}

	userObjectID := bson.ObjectIdHex(userID)

	var results []ChatTO
	err = database.GetDatabase().C(ChatsCollectionName).Find(bson.M{"$or": []bson.M{bson.M{"partner": userObjectID}, bson.M{"offering-creator-id": userObjectID}}}).All(&results)
	if err != nil {
		return nil, err
	}

	var chatObjects []*models.Chat

	for _, result := range results {
		chatPartner, err := GetUserShortByID(result.Partner)
		if err != nil {
			return nil, err
		}

		chatModel := &models.Chat{
			ID:         result.ID,
			Partner:    chatPartner,
			OfferingID: result.OfferingID,
		}

		chatObjects = append(chatObjects, chatModel)
	}

	return chatObjects, nil
}

//GetChatByID getting chat by id
func GetChatByID(chatID string) (*models.Chat, error) {
	var result ChatTO
	err := database.GetDatabase().C(ChatsCollectionName).FindId(bson.ObjectIdHex(chatID)).One(result)
	if err != nil {
		return nil, err
	}

	chatPartner, err := GetUserShortByID(result.Partner)
	if err != nil {
		return nil, err
	}

	chatModel := &models.Chat{
		ID:         result.ID,
		Partner:    chatPartner,
		OfferingID: result.OfferingID,
	}

	return chatModel, nil
}

//SaveChat save a Chat model
func SaveChat(principal models.Principal, chat models.Chat) error {
	chatTO := ChatTO{
		Partner:           chat.Partner.ID,
		OfferingID:        chat.OfferingID,
		ID:                chat.ID,
		OfferingCreatorID: principal.Userid,
	}

	if len(chatTO.ID) == 0 {
		chatTO.ID = string(bson.NewObjectId())
	}

	_, error := database.GetDatabase().C(ChatsCollectionName).UpsertId(chatTO.ID, chatTO)

	return error
}
