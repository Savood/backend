package dao

import (
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
)



//GetUserByID getting user by id
func GetUserByID(userID string) (*models.User, error) {
	var user *models.User

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//GetUserShortByID getting short user by id
func GetUserShortByID(userID string) (*models.UserShort, error) {
	var user *models.UserShort

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//SaveUser saving a user
func SaveUser(user *models.User) error {
	if len(user.ID) == 0 {
		user.ID = bson.NewObjectId()
	}

	_, error := database.GetDatabase().C(database.UsersCollectionName).UpsertId(user.ID, user)

	return error
}
