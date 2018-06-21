package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//OfferingsCollectionName collection of offerings in mongodb
const OfferingsCollectionName = "offerings"

//GetAllOfferingsByUserID Get all Offerings filtered by userId
func GetAllOfferingsByUserID(userID string) ([]*models.Offering, error) {
	var offerings []*models.Offering

	err := database.GetDatabase().C(OfferingsCollectionName).Find(bson.M{"creator-id": bson.ObjectIdHex(userID)}).All(offerings)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}

//GetOfferingByID get offering by id
func GetOfferingByID(offeringID string) (*models.Offering, error) {
	var offering *models.Offering

	err := database.GetDatabase().C(OfferingsCollectionName).FindId(bson.ObjectIdHex(offeringID)).One(offering)
	if err != nil {
		return nil, err
	}

	return offering, nil
}

//SaveOffering save an offering
func SaveOffering(offering *models.Offering) error {
	if len(offering.ID) == 0 {
		offering.ID = string(bson.NewObjectId())
	}

	_, error := database.GetDatabase().C(ChatsCollectionName).UpsertId(offering.ID, offering)

	return error
}