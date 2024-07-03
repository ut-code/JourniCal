package auth

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/env"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type userId = uint

var TokenCache = helper.NewMap[userId, *oauth2.Token]()
var TokenFromJSON *oauth2.Token

func init() {
	if env.USE_TOKEN_JSON {
		f, err := os.Open("./token.json")
		helper.PanicOn(err)
		err = json.NewDecoder(f).Decode(TokenFromJSON)
		helper.PanicOn(err)
		if env.STATIC_USER {
			SetToken(user.StaticUser, TokenFromJSON)
		}
	}
}

// entrypoint. use this if you don't know what you should use.
// this does not update user's token if it's expired, but don't care just generate it again
func TokenFromContext(db *gorm.DB, config *oauth2.Config, c echo.Context) (*oauth2.Token, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return nil, err
	}
	return RestoreUsersToken(config, u)
}

// Use this instead if you want to update user's token.
// This will update user's token as necessary.
func RestoreUsersToken(config *oauth2.Config, u *user.User) (*oauth2.Token, error) {
	cached, ok := TokenCache.Get(u.ID)
	if ok && cached.Valid() {
		return cached, nil
	}
	token := RawToken(u)
	// The token will auto-refresh as necessary.
	// src: https://pkg.go.dev/golang.org/x/oauth2?utm_source=godoc#Config.Client
	if IsEmpty(token) {
		return nil, errors.New("cannot restore user's token: user doesn't have a token")
	}
	config.Client(context.Background(), token)
	if !token.Valid() {
		return nil, errors.New("failed to revive token")
	}
	TokenCache.Set(u.ID, token)
	return token, nil
}

func ExchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return config.Exchange(ctx, code)
}

func SaveToken(db *gorm.DB, uid uint, token *oauth2.Token) error {
	var user *user.User
	err := db.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		return err
	}
	SetToken(user, token)
	return db.Save(user).Error
}

func RawToken(u *user.User) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		TokenType:    u.TokenType,
		Expiry:       u.TokenExpiry,
	}
}
func SetToken(u *user.User, token *oauth2.Token) {
	u.AccessToken = token.AccessToken
	u.RefreshToken = token.RefreshToken
	u.TokenType = token.TokenType
	u.TokenExpiry = token.Expiry
}

func IsEmpty(t *oauth2.Token) bool {
	// if the access token is not empty, other parts are probably not empty either (please don't just fill access token with random value and call it a "valid token")
	return t.AccessToken == ""
}
