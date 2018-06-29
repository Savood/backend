package user

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
)

//UsersCreateNewUserHandler Handler func: Create New User
func UsersCreateNewUserHandler(params users.CreateNewUserParams, principal *models.Principal) middleware.Responder {
	formats := strfmt.NewFormats()

	user := params.Body
	user.ID = bson.NewObjectId()
	user.Badges = []string{}

	err := user.Validate(formats)

	if err != nil {
		attribute := "idk!?"
		message := err.Error()
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	dao.SaveUser(user)
	return users.NewCreateNewUserOK().WithPayload(user)
}

//UsersDeleteUserByIDHandler Handler func: Delete a User
func UsersDeleteUserByIDHandler(params users.DeleteUserByIDParams, principal *models.Principal) middleware.Responder {
	if principal.Userid.Hex() != params.ID {
		attribute := "idk!?"
		message := "?!"
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	user, err := dao.GetUserByID(principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	user.Address = &models.Address{}
	user.Firstname = "Gelöschter"
	user.Lastname = "Nutzer"
	user.Phone = ""
	user.Description = ""
	user.Email = "deleted@chd.cx"
	user.Badges = []string{}
	dao.SaveUser(user)

	return users.NewDeleteUserByIDNoContent()
}

//UsersGetUserByIDHandler Handler func: get User by ID
func UsersGetUserByIDHandler(params users.GetUserByIDParams, principal *models.Principal) middleware.Responder {
	if params.ID == principal.Userid.Hex() {
		user, err := dao.GetUserByID(params.ID)
		if err != nil {
			attribute := "idk!?"
			message := "?!"
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		return users.NewGetUserByIDOK().WithPayload(user)
	}
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

//UsersUpdateUserByIDHandler Handler func: updating User by ID
func UsersUpdateUserByIDHandler(params users.UpdateUserByIDParams, principal *models.Principal) middleware.Responder {
	userNew := params.Body

	user, err := dao.GetUserByID(principal.Userid.Hex())
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return users.NewUpdateUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	user.Email = userNew.Email
	user.Description = userNew.Description
	user.Phone = userNew.Phone
	user.Lastname = userNew.Lastname
	user.Firstname = userNew.Firstname
	user.Address = userNew.Address

	err = dao.SaveUser(user)
	if err != nil {
		attribute := "idk!?"
		message := "?!"
		return users.NewUpdateUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return users.NewUpdateUserByIDNoContent()
}
