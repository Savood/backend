package user

import (
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/go-openapi/runtime/middleware"
	"git.dhbw.chd.cx/savood/backend/dao"
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

//UsersCreateNewUserHandler Handler func: Create New User
func UsersCreateNewUserHandler(params users.CreateNewUserParams, principal *models.Principal) middleware.Responder {
	formats := strfmt.NewFormats()

	user := params.Body
	user.ID = principal.Userid
	user.Email = strfmt.Email(principal.Email)
	user.Badges = []string{}

	err := user.Validate(formats)

	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	dao.SaveUser(user)
	return users.NewCreateNewUserOK().WithPayload(user)
}

//UsersDeleteUserByIDHandler Handler func: Delete a User
func UsersDeleteUserByIDHandler(params users.DeleteUserByIDParams, principal *models.Principal) middleware.Responder {
	if principal.Userid.Hex() != params.ID {
		attribute := "error"
		message := "forbidden"
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}
	user, err := dao.GetUserByID(principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
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
			attribute := "error"
			message := err.Error()
			return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
		}

		return users.NewGetUserByIDOK().WithPayload(user)
	}
	userShort, err := dao.GetUserShortByID(params.ID)
	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	userShortJSON, err := json.Marshal(userShort)
	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	user := models.User{}
	err = json.Unmarshal(userShortJSON, &user)
	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewGetUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return users.NewGetUserByIDOK().WithPayload(&user)
}

//UsersUpdateUserByIDHandler Handler func: updating User by ID
func UsersUpdateUserByIDHandler(params users.UpdateUserByIDParams, principal *models.Principal) middleware.Responder {
	userNew := params.Body

	user, err := dao.GetUserByID(principal.Userid.Hex())
	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewUpdateUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	if len(userNew.Email.String()) > 0 {
		user.Email = userNew.Email
	}

	if len(userNew.Description) > 0 {
		user.Description = userNew.Description
	}

	if len(userNew.Phone) > 0 {
		user.Phone = userNew.Phone
	}

	if len(userNew.Lastname) > 0 {
		user.Lastname = userNew.Lastname
	}

	if len(userNew.Firstname) > 0 {
		user.Firstname = userNew.Firstname
	}

	if userNew.Address != nil {
		user.Address = userNew.Address
	}

	err = dao.SaveUser(user)
	if err != nil {
		attribute := "error"
		message := err.Error()
		return users.NewUpdateUserByIDBadRequest().WithPayload(&models.InvalidParameterInput{Attribute: &attribute, Message: &message})
	}

	return users.NewUpdateUserByIDNoContent()
}
