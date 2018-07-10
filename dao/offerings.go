package dao

import (
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/database"
	"github.com/globalsign/mgo/bson"
	"errors"
)

//OfferingTO Transfer Object for Offering
type OfferingTO struct {
	ID bson.ObjectId `json:"_id"`

	CreatorID bson.ObjectId `json:"creatorid"`

	InnerOffering models.Offering
}

func injectCreatorShortAndID(offeringTO OfferingTO) (*OfferingTO, error) {
	creator, err := GetUserShortByID(offeringTO.CreatorID.Hex())
	if err != nil {
		return nil, err
	}

	offeringTO.InnerOffering.Creator = creator
	offeringTO.InnerOffering.ID = offeringTO.ID
	return &offeringTO, nil
}

func injectSlice(offeringTOs []OfferingTO) ([]*models.Offering, error) {

	var offerings []*models.Offering

	for _, oTO := range offeringTOs {
		o, err := injectCreatorShortAndID(oTO)
		if err != nil {
			return nil, err
		}
		offerings = append(offerings, &o.InnerOffering)
	}

	return offerings, nil
}

//GetAllOfferingsByUserID Get all Offerings filtered by userId
func GetAllOfferingsByUserID(userID string) ([]*models.Offering, error) {
	var offeringTOs []OfferingTO

	err := database.GetDatabase().C(database.OfferingsCollectionName).Find(bson.M{"creatorid": bson.ObjectIdHex(userID)}).All(&offeringTOs)
	if err != nil {
		return nil, err
	}

	offerings, err := injectSlice(offeringTOs)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}

//GetOfferingByID get offering by id
func GetOfferingByID(offeringID string) (*models.Offering, error) {
	var offeringTO *OfferingTO

	err := database.GetDatabase().C(database.OfferingsCollectionName).FindId(bson.ObjectIdHex(offeringID)).One(&offeringTO)
	if err != nil {
		return nil, err
	}

	offeringTO, err = injectCreatorShortAndID(*offeringTO)
	if err != nil {
		return nil, err
	}

	offering := offeringTO.InnerOffering

	return &offering, nil
}

//SaveOffering save an offering
func SaveOffering(offering *models.Offering) error {
	if len(offering.ID) == 0 {
		offering.ID = bson.NewObjectId()
	}

	if offering.Creator == nil {
		return errors.New("offering needs creator")
	}
	if offering.Creator.ID == "" {
		return errors.New("offering creator ID cannot be empty")
	}

	offeringTO := OfferingTO{
		ID:            offering.ID,
		CreatorID:     offering.Creator.ID,
		InnerOffering: *offering,
	}

	_, error := database.GetDatabase().C(database.OfferingsCollectionName).UpsertId(offeringTO.ID, offeringTO)

	return error
}

// DeleteOfferingByID deletes an offering by id
func DeleteOfferingByID(offeringID string) error {
	return database.GetDatabase().C(database.OfferingsCollectionName).RemoveId(bson.ObjectIdHex(offeringID))
}

// GetNearOfferings gets offerings with $nearSphereQuery
func GetNearOfferings(location models.OfferingLocation, maxDistance float64) ([]*models.Offering, error) {

	var offeringTOs []OfferingTO

	err := database.GetDatabase().C(database.OfferingsCollectionName).Find(
		map[string]interface{}{
			"location": map[string]interface{}{
				"$nearSphere": map[string]interface{}{
					"$geometry":    location,
					"$minDistance": 0,
					"$maxDistance": maxDistance,
				},
			}}).All(&offeringTOs)

	if err != nil {
		return nil, err
	}

	offerings, err := injectSlice(offeringTOs)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}
