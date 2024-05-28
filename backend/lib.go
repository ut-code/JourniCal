package main

import (
	"encoding/json"
	"log"
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
	cookie, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	err = json.Unmarshal([]byte(cookie.Value), token)
	return &token, nil
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
		TODO()
		// c.Redirect() // input url here
	}
}
