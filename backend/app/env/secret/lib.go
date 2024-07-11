package secret

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

var OAuth2Config *oauth2.Config // is nil if options.CredentialsSource == none.
var StaticToken *oauth2.Token   // is nil if options.TokenSource == db.
var AuthURL string              // is empty if OAuth2Config == nil.
var DSN string                  // may be empty.

func init() {
	DSN = env("DSN")
	loadCredentials()
	loadStaticToken()
	if OAuth2Config != nil {
		AuthURL = OAuth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	}
}

type innerCredential struct {
	ClientId                string   `json:"client_id"`
	ProjectId               string   `json:"project_id"`
	AuthUri                 string   `json:"auth_uri"`
	TokenUri                string   `json:"token_uri"`
	AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectUris            []string `json:"redirect_uris"`
}

type credential struct {
	Web innerCredential `json:"web"`
}

func loadCredentials() {
	switch options.CREDENTIALS_SOURCE {
	case options.CredentialsSourceNone:
		// skip
	case options.CredentialsSourceFile:
		OAuth2Config = readCredentialsFromFile("credentials.json")
	case options.CredentailsSourceEnv:
		OAuth2Config = readCredentialsFromEnv()
	default:
		log.Fatalln("unknown CREDENTIALS_SOURCE config found in app/env/secret:", options.CREDENTIALS_SOURCE)
	}
}

func loadStaticToken() {
	switch options.TOKEN_SOURCE {
	case options.TokenSourceDB:
		// do nothing
	case options.TokenSourceFile:
		StaticToken = readTokenFromFile("./token.json")
	case options.TokenSourceEnv:
		StaticToken = readTokenFromEnv()
	default:
		log.Fatalln("unknown TOKEN_SOURCE in app/env/secret", options.TOKEN_SOURCE)
	}
}

func readTokenFromFile(filename string) *oauth2.Token {
	t := new(oauth2.Token)
	f, err := os.Open(filename)
	helper.ErrorLog(err)
	err = json.NewDecoder(f).Decode(t)
	helper.ErrorLog(err)
	return t
}

func readTokenFromEnv() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  some("TOKEN_ACCESS_TOKEN"),
		TokenType:    some("TOKEN_TOKEN_TYPE"),
		RefreshToken: some("TOKEN_REFRESH_TOKEN"),
		Expiry:       validTime("TOKEN_EXPIRY"),
	}
}

func readCredentialsFromFile(path string) *oauth2.Config {
	bytes, err := os.ReadFile(path)
	helper.ErrorLog(err, "Failed reading credentials.json")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	helper.ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}

func readCredentialsFromEnv() *oauth2.Config {
	credentialData := credential{Web: innerCredential{
		ClientId:                some("CREDENTIAL_CLIENT_ID"),
		ProjectId:               some("CREDENTIAL_PROJECT_ID"),
		AuthUri:                 validURL("CREDENTIAL_AUTH_URI"),
		TokenUri:                validURL("CREDENTIAL_TOKEN_URI"),
		AuthProviderX509CertUrl: validURL("CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL"),
		ClientSecret:            some("CREDENTIAL_CLIENT_SECRET"),
		RedirectUris:            []string{validURL("CREDENTIAL_REDIRECT_URIS")},
	}}
	bytes, err := json.Marshal(credentialData)
	helper.ErrorLog(err, "Failed struct credentials")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	helper.ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}

func env(name string) string {
	return os.Getenv(name)
}

// this functions asserts that env is not be empty.
// if os.Getenv(name) == "", it will panic.
//
// Expected use case:
//
// env := some("SOME_ENV") // panics if $SOME_ENV == ""
func some(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Fatalln("Empty environment variable:", name)
	}
	return env
}

// trim ' ', '\n', and '"' from around $name and returns it,
// as these are probably not expected.
//
// expected use case:
//
// env := trimmed("SOME_ENV") // panics if $SOME_ENV is "", " ", "\n ", etc...
//
// fmt.Println(env) // " foo  " becomes "foo", "\" bar \n\n  \n \"" becomes "bar"
func trimmed(name string) string {
	env := os.Getenv(name)
	env = strings.Trim(env, " \n\"")
	if env == "" {
		log.Fatalln("Empty environment variable:", name)
	}
	return env
}

// assert $name is a valid URL.
//
// expected use case:
//
// url := validURL("WEB_ORIGIN") // panics if $WEB_ORIGIN is "somerandomstring". doesn't panic if $WEB_ORIGIN is "http://localhost:3000/".
func validURL(name string) string {
	env := trimmed(name)
	if _, err := url.Parse(env); err != nil {
		log.Fatalln("Invalid url: ", name, env)
	}
	return env
}

// assert that $name fulfills RFC3339.
// otherwise it will panic.
//
// expected use case:
//
// var time time.Time = validTime("LAST_USED") // panics if $LAST_USED is not valid RFC3339
func validTime(name string) time.Time {
	env := trimmed(name)
	t, err := time.Parse(time.RFC3339, env)
	if err != nil {
		log.Fatalln("Invalid time formatting: ", name, err)
	}
	return t
}
