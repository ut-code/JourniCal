package user

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/env/secret"
	"github.com/ut-code/JourniCal/backend/pkg/cookie"
	"github.com/ut-code/JourniCal/backend/pkg/hash"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var StaticUser *User

func init() {
	var token *oauth2.Token
	if options.STATIC_TOKEN {
		token = secret.StaticToken
	}
	if options.STATIC_USER {
		u := New("test user", "test password", "random", "value", token)
		StaticUser = &u
	}
}

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

func (u *User) Save(c echo.Context) {
	s := u.SessionUser()
	SaveSessionUserToCookie(c, &s)
}

// ---------------- minimum lifetime, clear interval

var ucache = cache.New(20*time.Minute, 10*time.Minute)

// don't just read from cookie username, instead use this.
func FromEchoContext(db *gorm.DB, c echo.Context) (*User, error) {
	// read from context data
	// we don't need to save it here, because it is done at middleware
	cu, ok := c.Get("user").(User)
	if ok {
		return &cu, nil
	}

	if options.STATIC_USER {
		return StaticUser, nil
	}
	var u *User
	su, err := SessionUserFromCookie(c)
	if err != nil {
		return nil, err
	}

	// this doesn't need to be a hash, but I couldn't find a better method
	key := hash.SHA256(su).Base64()
	ptr, ok := ucache.Get(key)
	if !ok {
		goto access_db
	}
	u, ok = ptr.(*User)
	if !ok {
		log.Printf("unexpected casting failure at user.FromEchoContext()")
		goto access_db
	}
	return u, nil

access_db:
	u, err = FindUserFromSession(db, *su)
	if err != nil {
		return nil, err
	}
	ucache.Set(key, u, 0) // use default expiration
	return u, nil
}
