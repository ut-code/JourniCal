package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	cache "github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) echo.MiddlewareFunc {
	// todo: use some sort of cacher
	// s.t. we don't have to request database every request
	// purpose: reduce db query.
	c := cache.New(5*time.Minute, 10*time.Minute)
	_ = c
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// insert cache here and escape early if cache is found.
			u, err := FromEchoContext(db, c)
			if err != nil {
				return c.String(http.StatusUnauthorized, "you are either not logged in or invalid user")
			}
			if u == nil {
				return c.String(http.StatusInternalServerError, "user.FromEchoContext returned (nil, nil), which is not expected."+
					"if this ever happens, I might rewrite the entire server in Rust.")
			}
			c.Set("user", u)
			// save this to the cache too
			return next(c)
		}
	}
}
