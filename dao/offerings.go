package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

//OfferingsCollectionName collection of offerings in mongodb
const OfferingsCollectionName = "offerings"

//OfferingDAO DAO for Offering
type OfferingDAO struct {
}

//GetAllByUserID Get all Offerings filtered by userId
func (dao OfferingDAO) GetAllByUserID(userID string) ([]*models.Offering, error) {
	var offerings []*models.Offering

	err := database.GetDatabase().C(OfferingsCollectionName).Find(bson.M{"creator-id": userID}).All(offerings)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}
