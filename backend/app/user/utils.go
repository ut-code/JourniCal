package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inSeconds int

const (
	hour  inSeconds = 60 * 60
	day             = 24 * hour
	week            = 7 * day
	month           = 30 * day
	year            = 365 * day
)

func SessionFromContext(c echo.Context) (string, error) {
	return Cookie(c, "session")
}

func Cookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// session user this function returns is not always valid.
// attackers can send whatever and this won't detect.
func SessionUserFromCookie(c echo.Context) (*SessionUser, error) {
	username, err := Cookie(c, "username")
	if err != nil {
		return nil, err
	}
	session, err := Cookie(c, "session")
	if err != nil {
		return nil, err
	}
	idString, err := Cookie(c, "id")
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	return &SessionUser{
		ID:       id,
		Username: username,
		Session:  session,
	}, nil
}

func FromEchoContext(db *gorm.DB, c echo.Context) (*User, error) {
	su, err := SessionUserFromCookie(c)
	if err != nil {
		return nil, err
	}
	u, err := FindUserFromSession(db, *su)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func AddSessionCookieOnContext(c echo.Context, session string) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    session,
		MaxAge:   1 * year,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	c.SetCookie(&cookie)
}
