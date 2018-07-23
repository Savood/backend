package dao

import (
	"github.com/go-openapi/strfmt"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"time"
)

//MessageTO Transfer Object for Chat
type MessageTO struct {
	ID bson.ObjectId `json:"_id"`

	ChatID bson.ObjectId `json:"chat-id"`

	Content string `json:"content"`

	From bson.ObjectId `json:"partner"`

	Time bson.MongoTimestamp `json:"time"`
}

//GetAllMessagesByChatID getting all messages in a chat
func GetAllMessagesByChatID(chatID string) ([]*models.Message, error) {
	var results []MessageTO

	err := database.GetDatabase().C(database.MessagesCollectionName).Find(bson.M{"chatid": bson.ObjectIdHex(chatID)}).All(&results)
	if err != nil {
		return nil, err
	}

	var messageObjects []*models.Message

	for _, result := range results {
		userFrom, err := GetUserShortByID(result.From.Hex())
		if err != nil {
			return nil, err
		}

		messageModel := &models.Message{
			Content: result.Content,
			From:    userFrom,
			Time:    strfmt.DateTime(time.Unix(int64(result.Time), 0)),
		}

		messageObjects = append(messageObjects, messageModel)
	}

	return messageObjects, nil
}

//SaveMessage save a message
func SaveMessage(chat models.Chat, message models.Message) error {
	messageTO := MessageTO{
		From:    message.From.ID,
		ID:      bson.NewObjectId(),
		Time:    bson.MongoTimestamp(time.Time(message.Time).Unix()),
		Content: message.Content,
		ChatID:  chat.ID,
	}

	error := database.GetDatabase().C(database.MessagesCollectionName).Insert(messageTO)
	return error
}
