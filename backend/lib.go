package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

func srvFromContext(c echo.Context) (*calendar.Service, error) {
	token, err := readToken(c)
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	c := newOAuthClient(ctx, config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(c))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	return nil, error("not implemented yet")
}
