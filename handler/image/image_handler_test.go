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

// Offerings

func TestGetOfferingsIDImageHandler1(t *testing.T) {
	response := GetOfferingsIDImageJpegHandler(operations.GetOfferingsIDImageJpegParams{ID: "500000000000000000000003"}, testPrincipal)

	assert.IsType(t, &operations.GetOfferingsIDImageJpegOK{}, response)

}

func TestGetOfferingsIDImageHandler2(t *testing.T) {
	height := float64(2000)
	width := float64(50)

	response := GetOfferingsIDImageJpegHandler(operations.GetOfferingsIDImageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetOfferingsIDImageJpegOK{}, response)

}

func TestGetOfferingsIDImageHandler3(t *testing.T) {
	height := float64(2000)
	width := float64(-50)

	response := GetOfferingsIDImageJpegHandler(operations.GetOfferingsIDImageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetOfferingsIDImageJpegOK{}, response)
}

func TestGetOfferingsIDImageHandler4(t *testing.T) {
	response := GetOfferingsIDImageJpegHandler(operations.GetOfferingsIDImageJpegParams{ID: "404040404040404040404040"}, testPrincipal)

	assert.IsType(t, &operations.GetOfferingsIDImageJpegNotFound{}, response)

}

// Users

func TestGetUsersIDImageHandler1(t *testing.T) {
	response := GetUsersIDImageJpegHandler(operations.GetUsersIDImageJpegParams{ID: "500000000000000000000003"}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDImageJpegOK{}, response)

}

func TestGetUsersIDImageHandler2(t *testing.T) {
	height := float64(2000)
	width := float64(50)

	response := GetUsersIDImageJpegHandler(operations.GetUsersIDImageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDImageJpegOK{}, response)

}

func TestGetUsersIDImageHandler3(t *testing.T) {
	height := float64(2000)
	width := float64(-50)

	response := GetUsersIDBackgroundimageJpegHandler(operations.GetUsersIDBackgroundimageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDImageJpegOK{}, response)
}

func TestGetUsersIDImageHandler4(t *testing.T) {
	response := GetUsersIDImageJpegHandler(operations.GetUsersIDImageJpegParams{ID: "404040404040404040404040"}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDImageJpegNotFound{}, response)

}

func TestPostUsersIDImageHandler1(t *testing.T) {

	img, err := os.Open("testimage.jpg")

	assert.NoError(t, err)

	response := PostUsersIDImageJpegHandler(operations.PostUsersIDImageJpegParams{ID: "500000000000000000000003", Upfile: img}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDImageJpegForbidden{}, response)
}

func TestPostUsersIDImageHandler2(t *testing.T) {

	img, err := os.Open("testimage.jpg")

	assert.NoError(t, err)

	response := PostUsersIDImageJpegHandler(operations.PostUsersIDImageJpegParams{ID: "5b32d488129072313591c682", Upfile: img}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDImageJpegNoContent{}, response)
}

func TestPostUsersIDImageHandler3(t *testing.T) {

	response := PostUsersIDImageJpegHandler(operations.PostUsersIDImageJpegParams{ID: "5b32d488129072313591c682"}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDImageJpegBadRequest{}, response)
}

// User Background

func TestGetUsersIDBackgroundImageHandler1(t *testing.T) {
	response := GetUsersIDBackgroundimageJpegHandler(operations.GetUsersIDBackgroundimageJpegParams{ID: "500000000000000000000003"}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDBackgroundimageJpegOK{}, response)

}

func TestGetUsersIDBackgroundImageHandler2(t *testing.T) {
	height := float64(2000)
	width := float64(50)

	response := GetUsersIDBackgroundimageJpegHandler(operations.GetUsersIDBackgroundimageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDBackgroundimageJpegOK{}, response)

}

func TestGetUsersIDBackgroundImageHandler3(t *testing.T) {
	height := float64(2000)
	width := float64(-50)

	response := GetUsersIDBackgroundimageJpegHandler(operations.GetUsersIDBackgroundimageJpegParams{ID: "500000000000000000000003", Height: &height, Width: &width}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDBackgroundimageJpegOK{}, response)
}

func TestGetUsersIDBackgroundImageHandler4(t *testing.T) {
	response := GetUsersIDBackgroundimageJpegHandler(operations.GetUsersIDBackgroundimageJpegParams{ID: "404040404040404040404040"}, testPrincipal)

	assert.IsType(t, &operations.GetUsersIDBackgroundimageJpegNotFound{}, response)

}

func TestPostUsersIDBackgroundImageHandler1(t *testing.T) {

	img, err := os.Open("testimage.jpg")

	assert.NoError(t, err)

	response := PostUsersIDBackgroundimageJpegHandler(operations.PostUsersIDBackgroundimageJpegParams{ID: "500000000000000000000003", Upfile: img}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDBackgroundimageJpegForbidden{}, response)
}

func TestPostUsersIDBackgroundImageHandler2(t *testing.T) {

	img, err := os.Open("testimage.jpg")

	assert.NoError(t, err)

	response := PostUsersIDBackgroundimageJpegHandler(operations.PostUsersIDBackgroundimageJpegParams{ID: "5b32d488129072313591c682", Upfile: img}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDBackgroundimageJpegNoContent{}, response)
}

func TestPostUsersIDBackgroundImageHandler3(t *testing.T) {

	response := PostUsersIDImageJpegHandler(operations.PostUsersIDImageJpegParams{ID: "5b32d488129072313591c682"}, testPrincipal)

	assert.IsType(t, operations.PostUsersIDBackgroundimageJpegBadRequest{}, response)
}