package image

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/models"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"os"
	"git.dhbw.chd.cx/savood/backend/database/image"
)

var testPrincipal = &models.Principal{Email: "test@test.com", Userid: bson.ObjectId("5b32d488129072313591c682")}

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}

	img, err := os.Open("testimage.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = image.UploadImage("user_avatar_500000000000000000000003.jpg", img)
	if err != nil {
		log.Fatal(err)
		return
	}

	img, err = os.Open("testimage.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = image.UploadImage("offering_avatar_500000000000000000000003.jpg", img)
	if err != nil {
		log.Fatal(err)
		return
	}

	os.Exit(m.Run())
}

func TestGetOfferingsIDImageHandler1(t *testing.T) {
	response := GetOfferingsIDImageHandler(operations.GetOfferingsIDImageParams{ID: "500000000000000000000003"}, testPrincipal)

	assert.IsType(t, operations.GetOfferingsIDImageOK{}, response)

}

func TestGetOfferingsIDImageHandler2(t *testing.T) {
	height := float64(2000)
	width := float64(50)

	response := GetOfferingsIDImageHandler(operations.GetOfferingsIDImageParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, operations.GetOfferingsIDImageOK{}, response)

}

func TestGetOfferingsIDImageHandler3(t *testing.T) {
	height := float64(2000)
	width := float64(-50)

	response := GetOfferingsIDImageHandler(operations.GetOfferingsIDImageParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, operations.GetOfferingsIDImageOK{}, response)
}

func TestGetOfferingsIDImageHandler4(t *testing.T) {
	response := GetOfferingsIDImageHandler(operations.GetOfferingsIDImageParams{ID: "404040404040404040404040"}, testPrincipal)

	assert.IsType(t, operations.GetOfferingsIDImageNotFound{}, response)

}

func TestGetUsersIDImageHandler1(t *testing.T) {
	response := GetUsersIDImageHandler(operations.GetUsersIDImageParams{ID: "500000000000000000000003"}, testPrincipal)

	assert.IsType(t, operations.GetUsersIDImageOK{}, response)

}

func TestGetUsersIDImageHandler2(t *testing.T) {
	height := float64(2000)
	width := float64(50)

	response := GetUsersIDImageHandler(operations.GetUsersIDImageParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, operations.GetUsersIDImageOK{}, response)

}

func TestGetUsersIDImageHandler3(t *testing.T) {
	height := float64(2000)
	width := float64(-50)

	response := GetUsersIDImageHandler(operations.GetUsersIDImageParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, operations.GetUsersIDImageOK{}, response)
}

func TestGetUsersIDImageHandler4(t *testing.T) {
	response := GetUsersIDImageHandler(operations.GetUsersIDImageParams{ID: "404040404040404040404040"}, testPrincipal)

	assert.IsType(t, operations.GetUsersIDImageNotFound{}, response)

}
