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
	"golang.org/x/oauth2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
---
README
---
you need to have token.secret at somewhere to test this.
TODO: write how to obtain it
*/

var db *gorm.DB
var config *oauth2.Config

func init() {
	os.Remove("./test.db")
	var err error
	db, err = gorm.Open(sqlite.Open("./test.db"))
	helper.PanicOn(err)
	err = db.AutoMigrate(&user.User{})
	helper.PanicOn(err)

	config = calendar.Config
}

func TestBasicFunctionality(t *testing.T) {
	assert := assert.New(t)

	// FIXME: get token from somewhere
	token, err := readTestingToken()
	helper.PanicOn(err)

	// test RestoreUsersToken
	seed := "random seed value"
	u, err := user.CreateUser(db, "username", "password_hashed_at_frontend", seed, seed, token)
	assert.Nil(err)

	tok, err := auth.RestoreUsersToken(config, u)
	assert.Nil(err)          // this is not nil for now
	assert.True(tok.Valid()) // this is not valid now

	// test TokenFromContext is skipped because I can't provide echo.Context, and the only thing it uses echo.Context for is to get user from it
}

func readTestingToken() (*oauth2.Token, error) {
	f, err := os.Open("./token.secret")
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
