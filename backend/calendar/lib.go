package calendar

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func SrvFromContext(c echo.Context) (*calendar.Service, error) {
	token, err := ReadToken(c)
	if err != nil {
		c.Redirect(http.StatusFound, AuthURL)
		return nil, err
	}
	client := cfg.Client(ctx, token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		c.String(500, "Internal Error: calendar.NewService failed")
		return nil, err
	}
	return srv, nil
}

func ReadToken(c echo.Context) (*oauth2.Token, error) {
	// maybe this should be stored in a database?
	// reason: getting token from same code twice doesn't seem to be possible
	var code string
	{
		cookie, err := c.Cookie("code")
		if err != nil {
			return nil, err
		}
		code, err = url.QueryUnescape(cookie.Value)
		if err != nil {
			return nil, err
		}
	}

	// read from cache here
	if cachedToken, ok := TokenCache.Get(code); ok {
		return &cachedToken, nil
	}

	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		fmt.Println("Unable to retrieve token from web", err)
		return nil, err
	}
	TokenCache.Set(code, *token)

	return token, nil
}

func WriteAuthCodeToCookie(c echo.Context, code string) {
	MaxAge := (24 * time.Hour).Seconds() // about 1 day.
	c.SetCookie(&http.Cookie{
		Path:     "/",
		Name:     "code",
		Value:    url.QueryEscape(code),
		MaxAge:   int(MaxAge),
		HttpOnly: true, // reduces XSS risk via disallowing access from browser JS
		SameSite: http.SameSiteDefaultMode,
	})
}
