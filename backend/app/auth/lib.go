package auth

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type userId = uint

// TODO: use this.
var TokenCache = helper.NewMap[userId, *oauth2.Token]()

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
		return nil, errors.New("failed to revive token. blame google.oauth2, not us.")
	}
	TokenCache.Set(u.ID, token)
	return token, nil
}

func RawToken(u *user.User) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		TokenType:    u.TokenType,
		Expiry:       u.TokenExpiry,
	}
}

func IsEmpty(t *oauth2.Token) bool {
	// if the access token is not empty, other parts are probably not empty either (please don't just fill access token with random value and call it a "valid token")
	return t.AccessToken == ""
}
