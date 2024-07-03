package secret

import (
	"encoding/json"
	"os"

	"github.com/ut-code/JourniCal/backend/app/env"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

var OAuth2Config *oauth2.Config
var TokenFromJSON *oauth2.Token
var AuthURL string

func init() {
	if !env.NO_CREDENTIALS_REQUIRED {
		OAuth2Config = ReadCredentials("credentials.json")
		AuthURL = OAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	}
	if env.USE_TOKEN_JSON {
		TokenFromJSON = new(oauth2.Token)
		f, err := os.Open("./token.json")
		helper.ErrorLog(err)
		err = json.NewDecoder(f).Decode(TokenFromJSON)
		helper.ErrorLog(err)
	}
}

// this operation halts the app if there is no credentials.json found.
func ReadCredentials(path string) *oauth2.Config {
	bytes, err := os.ReadFile(path)
	helper.ErrorLog(err, "Failed reading credentials.json")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	helper.ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}
