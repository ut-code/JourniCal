package main

// auth.go is intended to be used for OAuth2 authentication/authorization
// and be used in calendar.go

import (
	"net/http"

	"context"

	"golang.org/x/oauth2"
)

type authUrl = string
type authCode = string

// tok can be nil. if tok is nil, this function will call getToken(url) and saveToken(tok).
func Authenticate(ctx context.Context, config *oauth2.Config, tok *oauth2.Token, getToken func(url authUrl) authCode, saveToken func(*oauth2.Token)) *http.Client {
	if tok == nil {
		tok = callGetToken(ctx, config, getToken)
		saveToken(tok)
	}
	return config.Client(ctx, tok)
}

func callGetToken(ctx context.Context, config *oauth2.Config, getToken func(url authUrl) authCode) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	code := getToken(authURL)

	tok, err := config.Exchange(ctx, code)
	ErrorLog(err, "Unable to retrieve token from web")

	return tok
}
