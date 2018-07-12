package dao

import (
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"errors"
)

type UserTO struct {
	ID bson.ObjectId `json:"_id"`

	InnerUser models.User `json:"innerUser"`

	Savoods []bson.ObjectId `json:"savoods"`
}

type UserShortTO struct {
	ID bson.ObjectId `json:"_id"`

	InnerUser models.UserShort `json:"innerUser"`
}

//GetUserByID getting user by id
func GetUserByID(userID string) (*models.User, error) {
	var uTO UserTO

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(bson.ObjectIdHex(userID)).One(&uTO)
	if err != nil {
		return nil, err
	}

	return &uTO.InnerUser, nil
}

//GetUserShortByID getting short user by id
func GetUserShortByID(userID string) (*models.UserShort, error) {
	var usTO UserShortTO

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(bson.ObjectIdHex(userID)).One(&usTO)
	if err != nil {
		return nil, err
	}

	return &usTO.InnerUser, nil
}

//SaveUser saving a user
func SaveUser(user *models.User) error {
	//TODO rework this to a non deleting update

	if len(user.ID) == 0 {
		user.ID = bson.NewObjectId()
	}

	uTO := UserTO{
		ID:        user.ID,
		InnerUser: *user,
	}

	_, err := database.GetDatabase().C(database.UsersCollectionName).UpsertId(uTO.ID, uTO)

	return err
}

func AddSavoodToUserID(userID, savoodID string) error {

	//TODO do this right

	log.Print(userID)
	log.Print(savoodID)

	return database.GetDatabase().C(database.UsersCollectionName).UpdateId(bson.ObjectIdHex(userID), struct {
		Savoods []bson.ObjectId `json:"savoods"`
	}{
		Savoods: []bson.ObjectId{bson.ObjectIdHex(savoodID)},
	})

}

func RemoveSavoodFromUserID(userID, savoodID string) error {

	//TODO do this right

	log.Print(userID)
	log.Print(savoodID)

	return errors.New("not implemented")

}

func GetSavoodsByUserID(userID bson.ObjectId) ([]bson.ObjectId, error) {
	var uTO UserTO

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(userID).One(&uTO)
	if err != nil {
		return nil, err
	}

	return uTO.Savoods, nil
}
