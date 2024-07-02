package router

import (
	"net/http"

	"github.com/ut-code/JourniCal/backend/app/auth"
	"github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group, db *gorm.DB, conf *oauth2.Config) {

	g.GET("/new", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, calendar.AuthURL)
	})

	g.GET("/check", func(c echo.Context) error {
		token, err := auth.TokenFromContext(db, conf, c)
		if err != nil {
			c.String(200, "You are not authenticated. access /auth/new to authenticate.")
		} else {
			c.String(200, "You are authenticated. the code is: "+helper.ToJSON(token))
		}
		return nil
	})

	g.GET("/code", func(c echo.Context) error {
		code := c.QueryParam("code")
		if code == "" {
			c.String(http.StatusBadRequest, "empty authorization code")
			// anyone can freely send a request to /auth/code so it's not an actual error.
			// ignoring this on purpose to prevent log flood.
			return nil
		}

		// assert: user has been defined before coming to this url
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			c.String(http.StatusUnauthorized, "you haven't registered user yet")
		}

		token, err := auth.ExchangeToken(calendar.Config, code)
		if err != nil {
			c.String(http.StatusBadRequest, "bad authorization code")
			return nil
		}
		err = auth.SaveToken(db, u, token)
		if err != nil {
			c.String(http.StatusInternalServerError, "failed to save your token: "+err.Error())
		}

		c.Redirect(http.StatusFound, "/")
		return nil
	})
}
