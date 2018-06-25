package user

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"encoding/json"
)

func UsersCreateNewUserHandler(params users.CreateNewUserParams, principal *models.Principal) middleware.Responder {

}

func UsersDeleteUserByIDHandler(params users.DeleteUserByIDParams, principal *models.Principal) middleware.Responder {

}

func UsersGetUserByIDHandler(params users.GetUserByIDParams, principal *models.Principal) middleware.Responder {
	if params.ID == principal.Userid.Hex() {
		user, err := dao.GetUserByID(params.ID)
		if err != nil {
			attribute := "idk!?"
			message := "?!"
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		return users.NewGetUserByIDOK().WithPayload(user)
	} else {
		userShort, err := dao.GetUserShortByID(params.ID)
		if err != nil {
			attribute := "idk!?"
			message := "?!"
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		userShortJSON, err := json.Marshal(userShort)
		if err != nil {
			attribute := "idk!?"
			message := "?!"
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		user := models.User{}
		err = json.Unmarshal(userShortJSON, &user)
		if err != nil {
			attribute := "idk!?"
			message := "?!"
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		return users.NewGetUserByIDOK().WithPayload(&user)
	}
}

func UsersUpdateUserByIDHandler(params users.UpdateUserByIDParams, principal *models.Principal) middleware.Responder {
	user := params.Body
	user.ID = principal.Userid
	err := dao.SaveUser(user)
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return users.NewUpdateUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return users.NewUpdateUserByIDNoContent()
}