package router

import (
	"net/http"

	"github.com/ut-code/JourniCal/backend/app/auth"
	"github.com/ut-code/JourniCal/backend/app/secret"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// TODO: make this function
func Root(g *echo.Group, db *gorm.DB) {

	g.GET("/", func(c echo.Context) error {
		_, err := auth.TokenFromContext(db, secret.OAuth2Config, c)
		if err != nil {
			c.Redirect(http.StatusFound, secret.AuthURL)
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
