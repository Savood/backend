package dao

import (
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo"
)

// UserTO Transfer Object for User
type UserTO struct {
	ID bson.ObjectId `json:"_id"`

	InnerUser models.User `json:"innerUser"`

	Savoods []bson.ObjectId `json:"savoods"`
}

// UserShortTO Transfer Object for User but with UserShort
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
	var uTO UserTO

	if len(user.ID) == 0 {
		user.ID = bson.NewObjectId()
	} else {
		err := database.GetDatabase().C(database.UsersCollectionName).FindId(user.ID).One(&uTO)
		if err != nil && err != mgo.ErrNotFound {
			return err
		}
	}

	uTO.ID = user.ID
	uTO.InnerUser = *user

	_, err := database.GetDatabase().C(database.UsersCollectionName).UpsertId(uTO.ID, uTO)

	return err
}

// AddSavoodToUserID adds a savood id to the user object
func AddSavoodToUserID(userID, savoodID string) error {
	return database.GetDatabase().C(database.UsersCollectionName).UpdateId(bson.ObjectIdHex(userID), bson.M{"$addToSet": bson.M{"savoods": bson.ObjectIdHex(savoodID)}})
}

// RemoveSavoodFromUserID removes a savood id from the user object
func RemoveSavoodFromUserID(userID, savoodID string) error {

	var uTO UserTO
	err := database.GetDatabase().C(database.UsersCollectionName).FindId(bson.ObjectIdHex(userID)).One(&uTO)
	if err != nil {
		return err
	}

	a := uTO.Savoods
	b := a[:0]
	for _, x := range a {
		if x.Hex() != savoodID {
			b = append(b, x)
		}
	}

	uTO.Savoods = b

	_, err = database.GetDatabase().C(database.UsersCollectionName).UpsertId(uTO.ID, uTO)

	return err

}

// GetSavoodsByUserID gets all IDs of savooded offerings
func GetSavoodsByUserID(userID bson.ObjectId) ([]bson.ObjectId, error) {
	var uTO UserTO

	err := database.GetDatabase().C(database.UsersCollectionName).FindId(userID).One(&uTO)
	if err != nil {
		return nil, err
	}

	return uTO.Savoods, nil
}
