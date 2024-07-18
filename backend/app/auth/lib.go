// this package is for binding oauth2.Token and user together.
package auth

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/env/secret"
	"github.com/ut-code/JourniCal/backend/app/user"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

// entrypoint. use this if you don't know what you should use.
// this does not update user's token if it's expired, but don't care just generate it again
func TokenFromContext(db *gorm.DB, config *oauth2.Config, c echo.Context) (*oauth2.Token, error) {
	if options.STATIC_TOKEN {
		return secret.StaticToken, nil
	}
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return nil, err
	}
	return Token(u)
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

func Token(u *user.User) (*oauth2.Token, error) {
	// the token will auto-refresh as necessary.
	// src: https://pkg.go.dev/golang.org/x/oauth2?utm_source=godoc#Config.Client
	token := &oauth2.Token{
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		TokenType:    u.TokenType,
		Expiry:       u.TokenExpiry,
	}
	if IsEmpty(token) {
		return nil, errors.New("cannot restore user's token: user doesn't have a token")
	}
	return token, nil
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
