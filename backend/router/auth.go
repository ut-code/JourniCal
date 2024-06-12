package router

import (
	"github.com/ut-code/JourniCal/backend/calendar"
	"github.com/ut-code/JourniCal/backend/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group) {

	g.GET("/new", func(c echo.Context) error {
		c.Redirect(http.StatusFound, calendar.AuthURL)
		return nil
	})

	g.GET("/check", func(c echo.Context) error {
		token, err := calendar.ReadToken(c)
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

		calendar.WriteAuthCodeToCookie(c, code)

		c.Redirect(http.StatusFound, "/")
		return nil
	})

}
