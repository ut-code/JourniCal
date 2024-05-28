package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var cfg *oauth2.Config
var ctx context.Context
var authURL string

func getAuthNew(c echo.Context) error {
	c.Redirect(http.StatusFound, authURL)
	return nil
}

func postApiPing(c echo.Context) error {
	// read request body to bytes then stringify it
	// if you want to receive as JSON, there is a better way to do it:
	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	buf, err := io.ReadAll(c.Request().Body)
	ErrorLog(err)
	body := string(buf)
	fmt.Println(body)
	c.String(http.StatusOK, "pong!\n Your request body was: \n"+body+"\n")
	return nil
}

func getGet20EventsForward(c echo.Context) error {
	srv, err, no := srvFromContext(c)
	if err != nil {
		handleSrvInitializationError(c, no)
		return err
	}
	t, err := strconv.Atoi(c.Param("start_unix"))
	if err != nil {
		c.String(http.StatusBadRequest, "get-20-events-forward with invalid second path: not unix time (= number)")
	}
	evs := GetNEventsForward(srv, "primary", time.Unix(int64(t), 0), 20)
	return c.JSON(200, evs)
}

// TODO: make this function
func getRoot(c echo.Context) error {
	_, err := readToken(c)
	if err != nil {
		c.Redirect(http.StatusFound, authURL)
		return nil
	}
	c.File("./index.html")
	return nil
}

func getAuthCheck(c echo.Context) error {
	token, err := readToken(c)
	if err != nil {
		c.String(200, "You are not authenticated. access /auth/new to authenticate.")
	} else {
		c.String(200, "You are authenticated. the code is: "+toJSON(token))
	}
	return nil
}

func getAuthCode(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		c.String(http.StatusBadRequest, "empty authorization code")
		return nil // anyone can freely send a request to /auth/code so it's not an actual error
	}

	token, err := cfg.Exchange(ctx, code)
	ErrorLog(err, "Unable to retrieve token from web")
	b, err := json.Marshal(token)
	ErrorLog(err)

	const MaxAge = 12 * 30 * 24 * 60 * 60 // about 3 months.
	c.SetCookie(&http.Cookie{
		Path:     "/",
		Name:     "token",
		Value:    string(b),
		MaxAge:   MaxAge,
		HttpOnly: true, // reduces XSS risk via disallowing access from browser JS
	})
	c.Redirect(http.StatusFound, "/")
	return nil
}
