package router

import (
	"net/http"

	"github.com/ut-code/JourniCal/backend/app/auth"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// TODO: make this function
func Root(g *echo.Group, db *gorm.DB, authURL string, conf *oauth2.Config) {

	g.GET("/", func(c echo.Context) error {
		_, err := auth.TokenFromContext(db, conf, c)
		if err != nil {
			c.Redirect(http.StatusFound, authURL)
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
