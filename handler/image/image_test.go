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

	_, err := getImage("thisDoesNotExist.no")

	assert.Error(t, err)

}

func TestGetImage(t *testing.T) {

	testFile, err := os.Open("testimage.jpg")

	stats, _ := testFile.Stat()

	assert.NoError(t, err)

	if err == nil {
		err = uploadImage("testImage.jpg", testFile)

		assert.NoError(t, err)

		if err == nil {

			gridFile, err := getImage("testImage.jpg")

			g := gridFile.(*mgo.GridFile)

			assert.NoError(t, err)

			assert.Equal(t, stats.Size(), g.Size())
		}
	}

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

}
