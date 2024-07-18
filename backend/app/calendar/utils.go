package calendar

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/auth"
	"github.com/ut-code/JourniCal/backend/app/env/secret"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func SrvFromContext(db *gorm.DB, c echo.Context) (*calendar.Service, error) {
	conf := secret.OAuth2Config

	token, err := auth.TokenFromContext(db, conf, c)
	if err != nil {
		c.Redirect(http.StatusFound, secret.AuthURL)
		return nil, err
	}
	client := conf.Client(c.Request().Context(), token)
	srv, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
	if err != nil {
		c.String(500, "Internal Error: calendar.NewService failed")
		return nil, err
	}
	return srv, nil
}
