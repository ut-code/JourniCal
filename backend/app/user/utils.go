package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/pkg/cookie"
	"gorm.io/gorm"
)

// session user this function returns is not always valid.
// attackers can send whatever and this won't detect.
func SessionUserFromCookie(c echo.Context) (*SessionUser, error) {
	username, err := cookie.Get(c, "username")
	if err != nil {
		return nil, err
	}
	session, err := cookie.Get(c, "session")
	if err != nil {
		return nil, err
	}
	id, err := cookie.GetUint(c, "userid")
	if err != nil {
		return nil, err
	}
	return &SessionUser{
		ID:       id,
		Username: username,
		Session:  session,
	}, nil
}

func SaveSessionUserToCookie(c echo.Context, s *SessionUser) {
	cookie.SetUint(c, "userid", s.ID)
	cookie.Set(c, "username", s.Username)
	cookie.Set(c, "session", s.Session)
}

// don't just read from cookie username, instead use this.
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
