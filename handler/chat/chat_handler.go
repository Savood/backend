package chat

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/messages"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/go-openapi/strfmt"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
)

func MessagesGetAllChatsHandler(params messages.GetAllChatsParams, principal *models.Principal) middleware.Responder {
	chats, err := dao.GetAllChatsByUserID(principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return messages.NewGetAllChatsBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	return messages.NewGetAllChatsOK().WithPayload(chats)
}

func MessagesGetAllMessagesForChatHandler(params messages.GetAllMessagesForChatParams, principal *models.Principal) middleware.Responder {
	_, err := dao.GetChatByIDAndUserID(params.ChatID, principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return messages.NewGetAllMessagesForChatBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	msgs, err := dao.GetAllMessagesByChatID(params.ChatID)
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return messages.NewGetAllMessagesForChatBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	return messages.NewGetAllMessagesForChatOK().WithPayload(msgs)
}

func MessagesCreateNewMessageHandler(params messages.CreateNewMessageParams, principal *models.Principal) middleware.Responder {
	chat, err := dao.GetChatByIDAndUserID(params.ChatID, principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return messages.NewCreateNewMessageBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	message := params.Body
	message.From = &models.UserShort{
		ID: principal.Userid,
	}
	message.Time = strfmt.NewDateTime()

	dao.SaveMessage(*chat, *message)

	return messages.NewCreateNewMessageOK().WithPayload(message)
}

func MessagesGetAllChatsForOfferingHandler(params operations.GetAllChatsForOfferingParams, principal *models.Principal) middleware.Responder {
	chats, err := dao.GetAllChatsByOfferingAndUserID(params.ID, principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return operations.NewGetAllChatsForOfferingBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return operations.NewGetAllChatsForOfferingOK().WithPayload(chats)
}