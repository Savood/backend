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

	Location models.OfferingLocation `json:"location"`

	InnerOffering models.Offering `json:"innerOffering"`
}

func requestedByCount(savoodID bson.ObjectId) (int64, error) {
	n, err := database.GetDatabase().C(database.UsersCollectionName).Find(struct{ Savoods string `json:"savoods"` }{Savoods: savoodID.Hex()}).Count()

	return int64(n), err
}

func inject(offeringTO OfferingTO, principal *models.Principal) (*OfferingTO, error) {
	creator, err := GetUserShortByID(offeringTO.CreatorID.Hex())
	if err != nil {
		return nil, err
	}

	offeringTO.InnerOffering.Creator = creator
	offeringTO.InnerOffering.ID = offeringTO.ID
	offeringTO.InnerOffering.Location = &offeringTO.Location

	if principal != nil {
		savoods, err := GetSavoodsByUserID(principal.Userid)
		if err != nil {
			return &offeringTO, err
		}

		for _, id := range savoods {
			if offeringTO.ID == id {
				offeringTO.InnerOffering.Savooded = true
			}
		}
	}

	i, err := requestedByCount(offeringTO.ID)
	if err != nil {
		return &offeringTO, err
	}
	offeringTO.InnerOffering.RequestedBy = i

	return &offeringTO, nil
}

func injectSlice(offeringTOs []OfferingTO, principal *models.Principal) ([]*models.Offering, error) {

	var offerings []*models.Offering

	for _, oTO := range offeringTOs {
		o, err := inject(oTO, principal)
		if err != nil {
			return nil, err
		}
		offerings = append(offerings, &o.InnerOffering)
	}

	return offerings, nil
}

//GetAllOfferingsByUserID Get all Offerings filtered by userId
func GetAllOfferingsByUserID(userID string, principal *models.Principal) ([]*models.Offering, error) {
	var offeringTOs []OfferingTO

	err := database.GetDatabase().C(database.OfferingsCollectionName).Find(bson.M{"creatorid": bson.ObjectIdHex(userID)}).All(&offeringTOs)
	if err != nil {
		return nil, err
	}

	offerings, err := injectSlice(offeringTOs, principal)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}

//GetOfferingByID get offering by id
func GetOfferingByID(offeringID string, principal *models.Principal) (*models.Offering, error) {
	var offeringTO OfferingTO

	err := database.GetDatabase().C(database.OfferingsCollectionName).FindId(bson.ObjectIdHex(offeringID)).One(&offeringTO)
	if err != nil {
		return nil, err
	}

	oTO, err := inject(offeringTO, principal)
	if err != nil {
		return nil, err
	}

	offering := oTO.InnerOffering

	return &offering, nil
}

//SaveOffering save an offering
func SaveOffering(offering *models.Offering) error {

	if offering.Creator == nil {
		return errors.New("offering needs creator")
	}
	if offering.Creator.ID == "" {
		return errors.New("offering creator ID cannot be empty")
	}
	if offering.Location == nil {
		return errors.New("offering location cannot be empty")
	}

	if len(offering.ID) == 0 {
		offering.ID = bson.NewObjectId()
	}

	offeringTO := OfferingTO{
		ID:            offering.ID,
		CreatorID:     offering.Creator.ID,
		InnerOffering: *offering,
		Location:      *offering.Location,
	}

	_, err := database.GetDatabase().C(database.OfferingsCollectionName).UpsertId(offeringTO.ID, offeringTO)

	return err
}

// DeleteOfferingByID deletes an offering by id
func DeleteOfferingByID(offeringID string) error {
	return database.GetDatabase().C(database.OfferingsCollectionName).RemoveId(bson.ObjectIdHex(offeringID))
}

// GetNearOfferings gets offerings with $nearSphereQuery
func GetNearOfferings(location models.OfferingLocation, maxDistance float64, principal *models.Principal) ([]*models.Offering, error) {

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

	offerings, err := injectSlice(offeringTOs, principal)
	if err != nil {
		return nil, err
	}

	return offerings, nil
}
