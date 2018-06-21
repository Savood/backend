package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

const CHATS_COLLECTION_NAME = "chats"

type ChatDAO struct {
}

type ChatTO struct {
	ID string `json:"_id"`

	OfferingID []string `json:"offering-id"`

	Partner string `json:"partner"`
}

func (dao ChatDAO) GetAllByUserId(userId string) ([]*models.Chat, error) {
	offerings, err := OfferingDAO{}.GetAllByUserId(userId)
	if err != nil {
		return nil, err
	}

	var offeringIds []string

	for _, result := range offerings {
		offeringIds = append(offeringIds, result.ID)
	}

	var results []ChatTO
	err = database.GetDatabase().C(CHATS_COLLECTION_NAME).Find(bson.M{"$or": []bson.M{bson.M{"partner": userId}, bson.M{"offering-id": bson.M{"in": offeringIds}}}}).All(&results)
	if err != nil {
		return nil, err
	}

	var chatObjects []*models.Chat

	for _, result := range results {
		var chatPartner *models.UserShort

		err := database.GetDatabase().C(USERS_COLLECTION_NAME).FindId(result.Partner).One(chatPartner)
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

func (dao ChatDAO) SaveChat(chat models.Chat) error {
	chatTO := ChatTO{
		Partner: chat.Partner.ID,
		OfferingID: chat.OfferingID,
		ID: chat.ID,
	}

	if len(chatTO.ID) == 0 {
		chatTO.ID = string(bson.NewObjectId())
	}

	_, error := database.GetDatabase().C(CHATS_COLLECTION_NAME).UpsertId(chatTO.ID, chatTO)

	return error
}