package offerings

import (
	"testing"
	"git.dhbw.chd.cx/savood/backend/database"
	"log"
	"os"
	"github.com/globalsign/mgo/bson"
	"git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/dao"
	"github.com/go-openapi/strfmt"
)

var testPrincipal = &models.Principal{Email: "test@test.com", Userid: bson.ObjectIdHex("5b32d488129072313591c682")}
var forbiddenPrincipal = &models.Principal{Userid: bson.ObjectIdHex("404040404040404040404040")}

func TestMain(m *testing.M) {
	err := database.ConnectDatabase(nil, nil)
	if err != nil {
		log.Fatal("Failed DB connection")
		return
	}

	err = dao.SaveUser(&models.User{
		ID:    testPrincipal.Userid,
		Email: strfmt.Email(testPrincipal.Email),
	})
	if err != nil {
		log.Fatal("Failed User init")
		return
	}

	err = dao.SaveUser(&models.User{
		ID: forbiddenPrincipal.Userid,
	})
	if err != nil {
		log.Fatal("Failed User init")
		return
	}

	os.Exit(m.Run())
}
