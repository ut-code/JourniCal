package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	//"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func TODO(text ...string) {
	log.Fatal("TODO() was called: " + strings.Join(text, ", "))
}

func readToken(c echo.Context) (*oauth2.Token, error) {
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
	if cachedToken, ok := tokenCache.Get(code); ok {
		return &cachedToken, nil
	}

	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		fmt.Println("Unable to retrieve token from web", err)
		return nil, err
	}
	tokenCache.Set(code, *token)

	return token, nil
}

func writeAuthCode(c echo.Context, code string) {
	const MaxAge = 24 * 60 * 60 // about 1 day.
	c.SetCookie(&http.Cookie{
		Path:     "/",
		Name:     "code",
		Value:    url.QueryEscape(code),
		MaxAge:   MaxAge,
		HttpOnly: true, // reduces XSS risk via disallowing access from browser JS
	})
}

func toJSON[T any](v T) string {
	b, err := json.Marshal(v)
	ErrorLog(err)
	return string(b)
}

type errno int

const (
	OK = errno(iota)
	ERR_MISSING_TOKEN
	ERR_NEW_SERVICE_FAILED
)

func srvFromContext(c echo.Context) (*calendar.Service, error, errno) {
	token, err := readToken(c)
	if err != nil {
		return nil, err, ERR_MISSING_TOKEN
	}
	client := cfg.Client(ctx, token)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err, ERR_NEW_SERVICE_FAILED
	}
	return srv, nil, 0
}
func handleSrvInitializationError(c echo.Context, no errno) {
	if no == OK {
		// ok
		return
	}
	if no == ERR_MISSING_TOKEN {
		c.Redirect(http.StatusFound, authURL) // input url here
		return
	}
	if no == ERR_NEW_SERVICE_FAILED {
		c.String(500, "Internal Error: calendar.NewService failed")
		return
	}
}
