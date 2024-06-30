package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/random"
	"gorm.io/gorm"
)

func User(g *echo.Group, db *gorm.DB) {
	g.POST("/create/:username/:password/:frontendSeed", func(c echo.Context) error {
		username := c.Param("username")
		password := c.Param("password")
		frontendSeed := c.Param("frontendSeed")
		backendSeed := random.String(32)
		u, err := user.CreateUser(db, username, password, frontendSeed, backendSeed, nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, u)
	})
	g.GET("/debug/create/:username/:password", func(c echo.Context) error {
		username := c.Param("username")
		password := c.Param("password")
		u, err := user.CreateUser(db, username, password, "random", "value", nil)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	})
	g.GET("/debug/whoami", func(c echo.Context) error {
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusOK, "not found; error: "+err.Error())
		}
		return c.JSON(http.StatusOK, u)
	})
}
