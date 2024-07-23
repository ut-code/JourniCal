package calendar

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	cache "github.com/patrickmn/go-cache"
	"github.com/ut-code/JourniCal/backend/app/auth"
	"github.com/ut-code/JourniCal/backend/app/env/secret"
	"github.com/ut-code/JourniCal/backend/pkg/hash"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

var srvcache = cache.New(10*time.Minute, 15*time.Minute)

// handles sending responses if it errors
// reason: it's hard to handle at caller side
func SrvFromContext(db *gorm.DB, c echo.Context) (*calendar.Service, error) {

	token, err := auth.TokenFromContext(db, secret.OAuth2Config, c)
	if err != nil {
		err = c.Redirect(http.StatusFound, secret.AuthURL)
		return nil, err
	}

	key := hash.SHA256(token).Base64()
	var srv *calendar.Service

	any, ok := srvcache.Get(key)
	if !ok {
		goto regen
	}
	srv, ok = any.(*calendar.Service)
	if !ok {
		goto regen
	}
	return srv, nil

regen:

	srv, err = CreateService(c.Request().Context(), token)
	if err != nil {
		return nil, c.String(500, "Internal error: "+err.Error())
	}
	srvcache.Set(key, srv, 0)
	return srv, nil
}

func CreateService(ctx context.Context, token *oauth2.Token) (*calendar.Service, error) {
	client := secret.OAuth2Config.Client(ctx, token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		err := errors.New("calendar.NewService failed: " + err.Error())
		return nil, err
	}
	return srv, nil

}
