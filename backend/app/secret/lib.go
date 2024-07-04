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

type CredentialInnerStruct struct {
	ClientId                string    `json:"client_id"`
	ProjectId               string    `json:"project_id"`
	AuthUri                 string    `json:"auth_uri"`
	TokenUri                string    `json:"token_uri"`
	AuthProviderX509CertUrl string    `json:"auth_provider_x509_cert_url"`
	ClientSecret            string    `json:"client_secret"`
	RedirectUris            [1]string `json:"redirect_uris"`
}

type CredentialStruct struct {
	Web CredentialInnerStruct `json:"web"`
}

func init() {
	if !env.NO_CREDENTIALS_REQUIRED {
		if !env.CREDENTIAL_FROM_ENV {
			OAuth2Config = ReadCredentials("credentials.json")
		} else {
			OAuth2Config = ReadCredentialsFromEnv()
		}
		AuthURL = OAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	}
	if env.USE_TOKEN_JSON {
		if !env.TOKEN_FROM_ENV {
			TokenFromJSON = new(oauth2.Token)
			f, err := os.Open("./token.json")
			helper.ErrorLog(err)
			err = json.NewDecoder(f).Decode(TokenFromJSON)
			helper.ErrorLog(err)
		} else {
			TokenFromJSON = &oauth2.Token{
				AccessToken:  env.TOKEN_ACCESS_TOKEN,
				TokenType:    env.TOKEN_TOKEN_TYPE,
				RefreshToken: env.TOKEN_REFRESH_TOKEN,
				Expiry:       env.TOKEN_EXPIRY,
			}
		}
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

func ReadCredentialsFromEnv() *oauth2.Config {
	credentialData := CredentialStruct{Web: CredentialInnerStruct{
		ClientId:                env.CREDENTIAL_CLIENT_ID,
		ProjectId:               env.CREDENTIAL_PROJECT_ID,
		AuthUri:                 env.CREDENTIAL_AUTH_URI,
		TokenUri:                env.CREDENTIAL_TOKEN_URI,
		AuthProviderX509CertUrl: env.CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL,
		ClientSecret:            env.CREDENTIAL_CLIENT_SECRET,
		RedirectUris:            env.CREDENTIAL_REDIRECT_URLS,
	}}
	bytes, err := json.Marshal(credentialData)
	helper.ErrorLog(err, "Failed struct credentials")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	helper.ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}
