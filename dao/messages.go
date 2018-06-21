package dao

import (
	"github.com/go-openapi/strfmt"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
)

//MessagesCollectionName collection of messages in mongodb
const MessagesCollectionName = "messages"

//MessageTO Transfer Object for Chat
type MessageTO struct {
	ID string `json:"_id"`

	ChatID string `json:"chat-id"`

	Content string `json:"content"`

	From string `json:"partner"`

	Time strfmt.DateTime `json:"time"`
}

//GetAllMessagesByChatID getting all messages in a chat
func GetAllMessagesByChatID(chatID string) ([]*models.Message, error) {
	var results []MessageTO

	err := database.GetDatabase().C(MessagesCollectionName).Find(bson.M{"chat-id": bson.ObjectIdHex(chatID)}).All(results)
	if err != nil {
		return nil, err
	}

	var messageObjects []*models.Message

	for _, result := range results {
		userFrom, err := GetUserShortByID(result.From)
		if err != nil {
			return nil, err
		}

		messageModel := &models.Message{
			Content: result.Content,
			From:    userFrom,
			Time:    result.Time,
		}

		messageObjects = append(messageObjects, messageModel)
	}

	return messageObjects, nil
}

//SaveMessage save a message
func SaveMessage(chatID models.Chat, message models.Message) error {
	messageTO := MessageTO{
		From:    message.From.ID,
		ID:      string(bson.NewObjectId()),
		Time:    message.Time,
		Content: message.Content,
		ChatID:  chatID.ID,
	}

	error := database.GetDatabase().C(MessagesCollectionName).Insert(messageTO)
	return error
}
