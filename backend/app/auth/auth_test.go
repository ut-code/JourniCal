package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/app/auth"
	cal "github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/env/secret"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	_ "github.com/ut-code/JourniCal/backend/pkg/tests/run-test-at-root"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
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
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"))
	helper.PanicOn(err)
	err = db.AutoMigrate(&user.User{})
	helper.PanicOn(err)

	config = secret.OAuth2Config
	authURL = secret.AuthURL
	token = secret.StaticToken
}

func TestBasicFunctionality(t *testing.T) {
	assert := assert.New(t)

	// test RestoreUsersToken
	seed := "random seed value"
	u, err := user.Create(db, "username", "password_hashed_at_frontend", seed, seed, token)
	assert.Nil(err)

	tok, err := auth.RestoreUsersToken(config, u)
	assert.Nil(err)
	assert.True(isValid(tok))

	// test TokenFromContext is skipped because I can't provide echo.Context, and the only thing it uses echo.Context for is to get user from it
}

func isValid(token *oauth2.Token) bool {
	client := config.Client(context.Background(), token)
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return false
	}
	evs, err := cal.GetNEventsForward(srv, "primary", time.Now(), 10)
	if err != nil {
		return false
	}
	return len(evs) != 0
}
