package image

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/globalsign/mgo"
	"bytes"
	"io"
)

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}
	os.Exit(m.Run())
}

func TestNotGettingImage(t *testing.T) {

	// try to get non existing file
	_, err := getImage("thisDoesNotExist.no")

	assert.Error(t, err)

}

func TestGetImage(t *testing.T) {

	// open local file
	testFile, err := os.Open("testimage.jpg")
	assert.NoError(t, err)

	// save stats about file
	stats, _ := testFile.Stat()

	if err != nil {
		return
	}

	// upload Image
	err = uploadImage("testImage.jpg", testFile)
	assert.NoError(t, err)

	if err != nil {
		return
	}

	// download Image
	gridFile, err := getImage("testImage.jpg")

	g := gridFile.(*mgo.GridFile)

	assert.NoError(t, err)

	// assert equal size
	assert.Equal(t, stats.Size(), g.Size())

	deleteImage("testImage.jpg")

}

func TestDeleteImage(t *testing.T) {
	// open Local file
	testFile, err := os.Open("testimage.jpg")
	assert.NoError(t, err)

	if err != nil {
		return
	}

	// upload image
	err = uploadImage("testImageDelete.jpg", testFile)
	assert.NoError(t, err)

	if err != nil {
		return
	}

	// delete image
	err = deleteImage("testImageDelete.jpg")
	assert.NoError(t, err)

	// assert no image is found
	_, err = getImage("testImageDelete.jpg")
	assert.Error(t, err)
}

func TestResizeImage(t *testing.T) {

	testFile, err := os.Open("testimage.jpg")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	err = uploadImage("testImage2.jpg", testFile)

	assert.NoError(t, err)

	if err != nil {
		return
	}

	testFileToResize, err := os.Open("testimage.jpg")
	assert.NoError(t, err)

	if err != nil {
		return
	}

	resizeLocal, err := resizeImage(testFileToResize, 5, 0)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	io.Copy(buf, resizeLocal)

	gridFile, err := getImage("testImage2.jpg")
	assert.NoError(t, err)

	resizeRemote, err := resizeImage(gridFile, 5, 0)
	assert.NoError(t, err)

	buf2 := new(bytes.Buffer)
	io.Copy(buf2, resizeRemote)

	assert.Equal(t, len(buf.Bytes()), len(buf2.Bytes()))

	deleteImage("testImage2.jpg")

}
