package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//Offering Transfer Object for Chat
type OfferingTO struct {
	ID bson.ObjectId `json:"_id"`

	CreatorID bson.ObjectId `json:"offeringcreatorid"`

}

//GetAllOfferingsByUserID Get all Offerings filtered by userId
func GetAllOfferingsByUserID(userID string) ([]*models.Offering, error) {
	var offerings []*models.Offering

	err := database.GetDatabase().C(database.OfferingsCollectionName).Find(bson.M{"creatorid": bson.ObjectIdHex(userID)}).All(&offerings)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}

//GetOfferingByID get offering by id
func GetOfferingByID(offeringID string) (*models.Offering, error) {
	var offering *models.Offering

	err := database.GetDatabase().C(database.OfferingsCollectionName).FindId(bson.ObjectIdHex(offeringID)).One(&offering)
	if err != nil {
		return nil, err
	}

	return offering, nil
}

//SaveOffering save an offering
func SaveOffering(offering *models.Offering) error {
	if len(offering.ID) == 0 {
		offering.ID = bson.NewObjectId()
	}

	_, error := database.GetDatabase().C(database.OfferingsCollectionName).UpsertId(offering.ID, offering)

	return error
}

// DeleteOfferingByID deletes an offering by id
func DeleteOfferingByID(offeringID string) error {
	return database.GetDatabase().C(database.OfferingsCollectionName).RemoveId(bson.ObjectIdHex(offeringID))
}


// GetNearOfferings gets offerings with $nearSphereQuery
func GetNearOfferings(location models.OfferingLocation, maxDistance float64) ([]*models.Offering, error) {

	var offerings []*models.Offering

	err := database.GetDatabase().C(database.OfferingsCollectionName).Find(
		map[string]interface{}{
			"location": map[string]interface{}{
				"$nearSphere": map[string]interface{}{
					"$geometry":    location,
					"$minDistance": 0,
					"$maxDistance": maxDistance,
				},
			}}).All(&offerings)

	if err != nil {
		return nil, err
	}

	return offerings, nil
}
