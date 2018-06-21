package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
)

const OFFERINGS_COLLECTION_NAME = "offerings"

type OfferingDAO struct {
}

func (dao OfferingDAO) GetAllByUserId(userId string) ([]*models.Offering, error) {
	var offerings []*models.Offering

	err := database.GetDatabase().C(OFFERINGS_COLLECTION_NAME).Find(bson.M{"creator-id": userId}).All(offerings)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}
