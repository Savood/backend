package offerings

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/globalsign/mgo"
)

// UnSavoodHandler delets a savood
func UnSavoodHandler(params operations.UnSavoodParams, principal *models.Principal) middleware.Responder {

	offering, err := dao.GetOfferingByID(params.OfferingID, principal)

	if err == mgo.ErrNotFound {
		return operations.NewUnSavoodNotFound()
	}

	if err != nil || offering == nil {
		str := err.Error()
		return operations.NewUnSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	if offering.Creator.ID == principal.Userid {
		str := "cannot savood own offering"
		return operations.NewUnSavoodForbidden().WithPayload(&models.ErrorModel{Message: &str})
	}

	if !offering.Savooded {
		return operations.NewUnSavoodOK().WithPayload(&models.SuccessObject{Success: true})
	}

	err = dao.RemoveSavoodFromUserID(principal.Userid.Hex(), params.OfferingID)
	if err != nil {
		str := err.Error()
		return operations.NewUnSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}

	chats, err := dao.GetAllChatsByOfferingAndUserID(params.OfferingID, principal.Userid.Hex())
	if err != nil {
		str := err.Error()
		return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
	}
	if len(chats) > 0 {

		b := chats[0].OfferingID[:0]
		for _, x := range chats[0].OfferingID {
			if x != offering.ID {
				b = append(b, x)
			}
		}

		chats[0].OfferingID = b

		err = dao.SaveChat(chats[0])

		if err != nil {
			str := err.Error()
			return operations.NewPlaceSavoodInternalServerError().WithPayload(&models.ErrorModel{Message: &str})
		}
	}

	return operations.NewUnSavoodOK().WithPayload(&models.SuccessObject{Success: true})
}
