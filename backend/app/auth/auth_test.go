package auth_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/app/auth"
	"github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	_ "github.com/ut-code/JourniCal/backend/pkg/tests/run-test-at-root"
	"golang.org/x/oauth2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
---
README
---
you need to have token.json at somewhere to test this.
TODO: write how to obtain it
*/

var db *gorm.DB
var config *oauth2.Config
var token *oauth2.Token
var authURL string

func init() {
	wd, err := os.Getwd()
	helper.PanicOn(err)
	fmt.Println("[auth test] working directory:", wd)
	os.Remove("./test.db")
	db, err = gorm.Open(sqlite.Open("./test.db"))
	helper.PanicOn(err)
	err = db.AutoMigrate(&user.User{})
	helper.PanicOn(err)

	config = calendar.ReadCredentials()
	authURL = config.AuthCodeURL("state-string", oauth2.AccessTypeOffline)
	token, err = readTestingToken()
	helper.PanicOn(err)
}

func TestBasicFunctionality(t *testing.T) {
	assert := assert.New(t)

	// test RestoreUsersToken
	seed := "random seed value"
	u, err := user.CreateUser(db, "username", "password_hashed_at_frontend", seed, seed, token)
	assert.Nil(err)

	tok, err := auth.RestoreUsersToken(config, u)
	assert.Nil(err)
	assert.True(tok.Valid())

	// test TokenFromContext is skipped because I can't provide echo.Context, and the only thing it uses echo.Context for is to get user from it
}

func readTestingToken() (*oauth2.Token, error) {
	f, err := os.Open("./token.json")
	if err != nil {
	}
	defer f.Close()
	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
