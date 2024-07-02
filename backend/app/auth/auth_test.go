package auth_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

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
you need to have token.json at somewhere to test this.
TODO: write how to obtain it
*/

var db *gorm.DB
var config *oauth2.Config
var token *oauth2.Token

func init() {
	os.Remove("./test.db")
	var err error
	db, err = gorm.Open(sqlite.Open("./test.db"))
	helper.PanicOn(err)
	err = db.AutoMigrate(&user.User{})
	helper.PanicOn(err)

	config = calendar.ReadCredentials()
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
		// there no token.json
		token, err := obtainTestingToken()
		if err != nil {
			return nil, err
		}
		f, err := os.Create("./token.json")
		if err != nil {
			return nil, err
		}
		json.NewEncoder(f).Encode(token)
		fmt.Println("run the test again.")
		os.Exit(1)
	}
	defer f.Close()
	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func obtainTestingToken() (*oauth2.Token, error) {
	fmt.Println("Go to this link and click ok: ", calendar.AuthURL)
	time.Sleep(10 * time.Second)
	handler := handler{ch: make(chan string)}
	go http.ListenAndServe(":3000", handler)

	code := <-handler.ch
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

type handler struct{ ch chan string }

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ch <- r.URL.Query().Get("code")
	fmt.Fprintf(w, "accepted")
	close(h.ch)
}
