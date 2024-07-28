package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u, err := FromEchoContext(db, c)
			if err != nil {
				return c.String(http.StatusUnauthorized, "you are either not logged in or invalid user")
			}
			if u == nil {
				return c.String(http.StatusInternalServerError, "user.FromEchoContext returned (nil, nil).")
			}
			c.Set("user", *u)
			return next(c)
		}
	}
}
