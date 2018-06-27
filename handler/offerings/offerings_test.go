package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"os"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
)

var testPrincipal = &models.Principal{Email: "test@test.com", Userid: bson.ObjectId("5b32d488129072313591c682")}


func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}
	os.Exit(m.Run())
}