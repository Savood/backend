package chat

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/messages"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/go-openapi/strfmt"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
)

//MessagesGetAllChatsHandler Handler func: getting all Chats a user participates in
func MessagesGetAllChatsHandler(params messages.GetAllChatsParams, principal *models.Principal) middleware.Responder {
	chats, err := dao.GetAllChatsByUserID(principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
		return messages.NewGetAllChatsBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	return messages.NewGetAllChatsOK().WithPayload(chats)
}

//MessagesGetAllMessagesForChatHandler Handler func: getting all messages in a Chat
func MessagesGetAllMessagesForChatHandler(params messages.GetAllMessagesForChatParams, principal *models.Principal) middleware.Responder {
	_, err := dao.GetChatByIDAndUserID(params.ChatID, principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
		return messages.NewGetAllMessagesForChatBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	msgs, err := dao.GetAllMessagesByChatID(params.ChatID)
	if err != nil {
		attribute := "error"
		message := err.Error()
		return messages.NewGetAllMessagesForChatBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	return messages.NewGetAllMessagesForChatOK().WithPayload(msgs)
}

//MessagesCreateNewMessageHandler Handler func: Create New Message for In Chat
func MessagesCreateNewMessageHandler(params messages.CreateNewMessageParams, principal *models.Principal) middleware.Responder {
	chat, err := dao.GetChatByIDAndUserID(params.ChatID, principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
		return messages.NewCreateNewMessageBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	message := params.Body
	message.From = &models.UserShort{
		ID: principal.Userid,
	}
	message.Time = strfmt.NewDateTime()

	dao.TouchChat(chat.ID.Hex())
	dao.SaveMessage(*chat, *message)

	return messages.NewCreateNewMessageOK().WithPayload(message)
}

//MessagesGetAllChatsForOfferingHandler Handler func: getting all chats for an offering
func MessagesGetAllChatsForOfferingHandler(params operations.GetAllChatsForOfferingParams, principal *models.Principal) middleware.Responder {
	chats, err := dao.GetAllChatsByOfferingAndUserID(params.ID, principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
		return operations.NewGetAllChatsForOfferingBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return operations.NewGetAllChatsForOfferingOK().WithPayload(chats)
}