package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// PlaceSavoodHandler places a savood
func PlaceSavoodHandler(params operations.PlaceSavoodParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.OfferingID, nil)

	if err == mgo.ErrNotFound {
		return operations.NewPlaceSavoodNotFound()
	}

	if err != nil || offering == nil {
		str := err.Error()
		return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.Creator.ID == principal.Userid {
		str := "cannot savood own offering"
		return operations.NewPlaceSavoodForbidden().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.Savooded {
		return operations.NewPlaceSavoodOK().WithPayload(&models.SuccessObject{Success: true})
	}

	err = dao.AddSavoodToUserID(principal.Userid.Hex(), params.OfferingID)
	if err != nil {
		str := err.Error()
		return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	chats, err := dao.GetAllChatsByOfferingAndUserID(params.OfferingID, principal.Userid.Hex())
	if err != nil {
		str := err.Error()
		return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}
	if len(chats) == 0 {
		err = dao.SaveChat(&models.Chat{OfferingID: []bson.ObjectId{offering.ID}, Partner: &models.UserShort{ID: principal.Userid}})
	} else {
		chats[0].OfferingID = append(chats[0].OfferingID, offering.ID)
		err = dao.SaveChat(chats[0])
	}
	if err != nil {
		str := err.Error()
		return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	return operations.NewPlaceSavoodOK().WithPayload(&models.SuccessObject{Success: true})

}
