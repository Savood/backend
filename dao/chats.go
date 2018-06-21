package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//ChatsCollectionName collection of chats in mongodb
const ChatsCollectionName = "chats"

//ChatDAO DAO for Chat
type ChatDAO struct {
}

//ChatTO Transfer Object for Chat
type ChatTO struct {
	ID string `json:"_id"`

	OfferingID []string `json:"offering-id"`

	Partner string `json:"partner"`
}

//GetAllByUserID Get All chats by user id (by offerings and partner)
func (dao ChatDAO) GetAllByUserID(userID string) ([]*models.Chat, error) {
	offerings, err := OfferingDAO{}.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var offeringIds []string

	for _, result := range offerings {
		offeringIds = append(offeringIds, result.ID)
	}

	var results []ChatTO
	err = database.GetDatabase().C(ChatsCollectionName).Find(bson.M{"$or": []bson.M{bson.M{"partner": userID}, bson.M{"offering-id": bson.M{"in": offeringIds}}}}).All(&results)
	if err != nil {
		return nil, err
	}

	var chatObjects []*models.Chat

	for _, result := range results {
		var chatPartner *models.UserShort

		err := database.GetDatabase().C(UsersCollectionName).FindId(result.Partner).One(chatPartner)
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

//SaveChat save a Chat model
func (dao ChatDAO) SaveChat(chat models.Chat) error {
	chatTO := ChatTO{
		Partner: chat.Partner.ID,
		OfferingID: chat.OfferingID,
		ID: chat.ID,
	}

	if len(chatTO.ID) == 0 {
		chatTO.ID = string(bson.NewObjectId())
	}

	_, error := database.GetDatabase().C(ChatsCollectionName).UpsertId(chatTO.ID, chatTO)

	return error
}