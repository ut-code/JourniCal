package router

import (
	"net/http"

	"github.com/ut-code/JourniCal/backend/app/calendar"

	"github.com/labstack/echo/v4"
)

// TODO: make this function
func Root(g *echo.Group) {

	g.GET("/", func(c echo.Context) error {
		_, err := calendar.ReadToken(c)
		if err != nil {
			c.Redirect(http.StatusFound, calendar.AuthURL)
			return nil
		}
		c.File("./index.html")
		return nil
	})
}

func Api(g *echo.Group) {
	g.GET("/ping", func(c echo.Context) error {
		c.String(http.StatusOK, "pong!")
		return nil
	})
}