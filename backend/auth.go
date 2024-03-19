package main

// auth.go is intended to be used for OAuth2 authentication/authorization
// and be used in calendar.go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"context"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type authCode = string

func SendToAuth(c echo.Context, config *oauth2.Config) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// http status code matters. http.StatusFound tells the browser to send a GET request to the URL, while others may require to send a POST request.
	c.Redirect(http.StatusFound, authURL)
}

func SaveToken(code authCode, token *oauth2.Token) {
	file := "./token.json"
	fmt.Printf("Saving token file to: %s\n", file)
	err := os.Remove("./token.json")
	ErrorLog(err, "Unable to delete previous file")
	f, err := os.Create(file)
	ErrorLog(err, "Unable to create file")

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// returns a valid token.
// automatically saves new token if necessary.
// returns error if saved token is expired && (code is invalid || couldn't connect to google service).
func ReadToken(ctx context.Context, code authCode, cfg *oauth2.Config) (*oauth2.Token, error) {
	var tok oauth2.Token

	b, err := os.ReadFile("./token.json") // TODO: read tok from database where code = $code
	if err != nil {
		// couldn't read from token.json
		goto exchange_token
	}
	err = json.Unmarshal(b, &tok)
	if err != nil {
		// invalid JSON; should rewrite token
		goto exchange_token
	}
	if !tok.Valid() {
		// invalid token. (expired?) ~~TODO: get another token via RefreshToken~~
		// Client constructor will automatically refresh this. so nothing to do here.
		return &tok, nil
	}
	return &tok, nil

exchange_token:
	p_tok, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	SaveToken(code, p_tok)
	return p_tok, nil
}
