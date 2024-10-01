package router

import (
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func Dev(g *echo.Group, db *gorm.DB) {
	g.GET("/delay/:time", func(c echo.Context) error {
		param := c.Param("time")
		t, err := strconv.Atoi(param)
		if err != nil {
			return c.String(400, "bad parameter encoding")
		}
		time.Sleep(time.Duration(t) * time.Second)
		return c.String(200, "slept well for "+param+" seconds.")
	})
}
