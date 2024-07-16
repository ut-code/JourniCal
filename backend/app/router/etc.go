package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Api(g *echo.Group) {
	g.GET("/ping", func(c echo.Context) error {
		c.String(http.StatusOK, "pong!")
		return nil
	})
}
