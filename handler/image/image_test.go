package image

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
)

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}
	m.Run()
}

