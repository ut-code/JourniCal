package calendar

import (
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/auth"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func SrvFromContext(db *gorm.DB, conf *oauth2.Config, c echo.Context) (*calendar.Service, error) {
	token, err := auth.TokenFromContext(db, conf, c)
	if err != nil {
		c.Redirect(http.StatusFound, AuthURL)
		return nil, err
	}
	client := Config.Client(ctx, token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		c.String(500, "Internal Error: calendar.NewService failed")
		return nil, err
	}
	return srv, nil
}

// unsafe; don't use this.
func WriteAuthCodeToCookie(c echo.Context, code string) {
	MaxAge := (24 * time.Hour).Seconds() // about 1 day.
	c.SetCookie(&http.Cookie{
		Path:     "/",
		Name:     "code",
		Value:    url.QueryEscape(code),
		MaxAge:   int(MaxAge),
		HttpOnly: true, // reduces XSS risk via disallowing access from browser JS
		SameSite: http.SameSiteDefaultMode,
	})
}
