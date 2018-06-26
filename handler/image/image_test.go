package image

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}
	os.Exit(m.Run())
}

func TestGetImage(t *testing.T) {
	// TODO
	assert.False(t,true)
}