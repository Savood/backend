package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//ChatTO Transfer Object for Chat
type ChatTO struct {
	ID bson.ObjectId `json:"_id"`

	OfferingID []bson.ObjectId `json:"offeringid"`

	Partner bson.ObjectId `json:"partner"`

	OfferingCreatorID bson.ObjectId `json:"offeringcreatorid"`

	LastUpdated bson.MongoTimestamp
}

//GetAllChatsByOfferingAndUserID Get All chats with this offering id and containing the right userid
func GetAllChatsByOfferingAndUserID(offeringID string, userID string) ([]*models.Chat, error) {
	userObjectID := bson.ObjectIdHex(userID)
	offeringObjectID := bson.ObjectIdHex(offeringID)

	var results []ChatTO
	err := database.GetDatabase().C(database.ChatsCollectionName).Find(bson.M{"$or": []bson.M{bson.M{"partner": userObjectID}, bson.M{"offeringcreatorid": userObjectID}}, "offeringid": bson.M{"$in": []bson.ObjectId{offeringObjectID}}}).Sort("-lastupdated").All(&results)
	if err != nil {
		return nil, err
	}

	var chatObjects []*models.Chat

	for _, result := range results {
		chatPartner, err := GetUserShortByID(result.Partner.Hex())
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

//GetAllChatsByUserID Get All chats by user id (by offerings and partner)
func GetAllChatsByUserID(userID string) ([]*models.Chat, error) {
	userObjectID := bson.ObjectIdHex(userID)

	var results []ChatTO
	err := database.GetDatabase().C(database.ChatsCollectionName).Find(bson.M{"$or": []bson.M{bson.M{"partner": userObjectID}, bson.M{"offering-creator-id": userObjectID}}}).Sort("-lastupdated").All(&results)
	if err != nil {
		return nil, err
	}

	var chatObjects []*models.Chat

	for _, result := range results {
		chatPartner, err := GetUserShortByID(result.Partner.Hex())
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

//GetChatByIDAndUserID getting chat by id
func GetChatByIDAndUserID(chatID string, userID string) (*models.Chat, error) {
	userObjectID := bson.ObjectIdHex(userID)

	var result ChatTO
	err := database.GetDatabase().C(database.ChatsCollectionName).Find(bson.M{"_id": bson.ObjectIdHex(chatID), "$or": []bson.M{bson.M{"partner": userObjectID}, bson.M{"offeringcreatorid": userObjectID}}}).One(&result)
	if err != nil {
		return nil, err
	}

	chatPartner, err := GetUserShortByID(result.Partner.Hex())
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

//GetChatByID getting chat by id
func GetChatByID(chatID string) (*models.Chat, error) {
	var result ChatTO
	err := database.GetDatabase().C(database.ChatsCollectionName).FindId(bson.ObjectIdHex(chatID)).One(&result)
	if err != nil {
		return nil, err
	}

	chatPartner, err := GetUserShortByID(result.Partner.Hex())
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

//UpdateChatRemoveOfferingID updates chat objects and removes the object id. Also removes not used chats.
func UpdateChatRemoveOfferingID(offeringID string) (error) {
	offeringObjectID := bson.ObjectIdHex(offeringID)

	_, err := database.GetDatabase().C(database.ChatsCollectionName).UpdateAll(bson.M{"offeringid": bson.M{"$in": []bson.ObjectId{offeringObjectID}}}, bson.M{"$pull": bson.M{"offeringid": offeringObjectID}})
	if err != nil {
		return err
	}

	_, err = database.GetDatabase().C(database.ChatsCollectionName).RemoveAll(bson.M{"offeringid": bson.M{"$size": 0}})
	if err != nil {
		return err
	}

	return nil
}

//TouchChat touching a chat last updated time
func TouchChat(chatID string) (error) {
	err := database.GetDatabase().C(database.ChatsCollectionName).UpdateId(bson.ObjectIdHex(chatID), bson.M{"$set": bson.M{"lastupdated": bson.MongoTimestamp(bson.Now().Unix())}})

	return err
}

//SaveChat save a Chat model
func SaveChat(chat *models.Chat) error {
	offeringID := chat.OfferingID[0]

	offering, err := GetOfferingByID(offeringID.Hex(), nil)
	if err != nil {
		return err
	}

	if len(chat.ID) == 0 {
		chat.ID = bson.NewObjectId()
	}

	chatTO := ChatTO{
		Partner:           chat.Partner.ID,
		OfferingID:        chat.OfferingID,
		ID:                chat.ID,
		OfferingCreatorID: offering.Creator.ID,
		LastUpdated:       bson.MongoTimestamp(bson.Now().Unix()),
	}

	_, error := database.GetDatabase().C(database.ChatsCollectionName).UpsertId(chatTO.ID, chatTO)

	return error
}
