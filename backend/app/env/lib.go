package env

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var USE_TOKEN_JSON = false
var STATIC_USER = false
var NO_CREDENTIALS_REQUIRED = false
var ECHO_SERVES_FRONTEND_TOO = false
var ENABLE_CORS = false
var CORS_ORIGIN string
var HALT_AFTER_SUCCESS = false
var DSN string
var CREDENTIAL_FROM_ENV = false
var TOKEN_FROM_ENV = false
var CREDENTIAL_CLIENT_ID string
var CREDENTIAL_PROJECT_ID string
var CREDENTIAL_AUTH_URI string
var CREDENTIAL_TOKEN_URI string
var CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL string
var CREDENTIAL_CLIENT_SECRET string
var CREDENTIAL_REDIRECT_URLS [1]string
var TOKEN_ACCESS_TOKEN string
var TOKEN_TOKEN_TYPE string
var TOKEN_REFRESH_TOKEN string
var TOKEN_EXPIRY time.Time

func EmptyCheck(variable string, message string) {
	if variable == "" {
		panic("empty environment variable: " + message)
	}
}

func init() {
	godotenv.Load()
	if os.Getenv("USE_TOKEN_JSON") == "true" {
		USE_TOKEN_JSON = true
	}
	if os.Getenv("STATIC_USER") == "true" {
		STATIC_USER = true
	}
	if os.Getenv("NO_CREDENTIALS_REQUIRED") == "true" {
		NO_CREDENTIALS_REQUIRED = true
	}
	if os.Getenv("ECHO_SERVES_FRONTEND_TOO") == "true" {
		ECHO_SERVES_FRONTEND_TOO = true
	}
	if corsOrigin := os.Getenv("CORS_ORIGIN"); corsOrigin != "" {
		ENABLE_CORS = true
		CORS_ORIGIN = corsOrigin
	}
	DSN = os.Getenv("DSN")
	if os.Getenv("CREDENTIAL_FROM_ENV") == "true" {
		CREDENTIAL_FROM_ENV = true
		CREDENTIAL_CLIENT_ID = os.Getenv("CREDENTIAL_CLIENT_ID")
		EmptyCheck(CREDENTIAL_CLIENT_ID, "CREDENTIAL_CLIENT_ID")
		CREDENTIAL_PROJECT_ID = os.Getenv("CREDENTIAL_PROJECT_ID")
		EmptyCheck(CREDENTIAL_PROJECT_ID, "CREDENTIAL_PROJECT_ID")
		CREDENTIAL_AUTH_URI = os.Getenv("CREDENTIAL_AUTH_URI")
		EmptyCheck(CREDENTIAL_AUTH_URI, "CREDENTIAL_AUTH_URI")
		CREDENTIAL_TOKEN_URI = os.Getenv("CREDENTIAL_TOKEN_URI")
		EmptyCheck(CREDENTIAL_TOKEN_URI, "CREDENTIAL_TOKEN_URI")
		CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL = os.Getenv("CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL")
		EmptyCheck(CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL, "CREDENTIAL_AUTH_PROVIDER_X509_CERT_URL")
		CREDENTIAL_CLIENT_SECRET = os.Getenv("CREDENTIAL_CLIENT_SECRET")
		EmptyCheck(CREDENTIAL_CLIENT_SECRET, "CREDENTIAL_CLIENT_SECRET")
		CREDENTIAL_REDIRECT_URLS[0] = os.Getenv("CREDENTIAL_REDIRECT_URLS")
		EmptyCheck(CREDENTIAL_REDIRECT_URLS[0], "CREDENTIAL_REDIRECT_URIS")
	}
	if os.Getenv("TOKEN_FROM_ENV") == "true" {
		TOKEN_FROM_ENV = true
		TOKEN_ACCESS_TOKEN = os.Getenv("TOKEN_ACCESS_TOKEN")
		EmptyCheck(TOKEN_ACCESS_TOKEN, "TOKEN_ACCESS_TOKEN")
		TOKEN_TOKEN_TYPE = os.Getenv("TOKEN_TOKEN_TYPE")
		EmptyCheck(TOKEN_TOKEN_TYPE, "TOKEN_TOKEN_TYPE")
		TOKEN_REFRESH_TOKEN = os.Getenv("TOKEN_REFRESH_TOKEN")
		EmptyCheck(TOKEN_REFRESH_TOKEN, "TOKEN_REFRESH_TOKEN")
		TOKEN_EXPIRY_STRING := os.Getenv("TOKEN_EXPIRY")
		TOKEN_EXPIRY_STRING = strings.Trim(TOKEN_EXPIRY_STRING, " \n")
		EmptyCheck(TOKEN_EXPIRY_STRING, "TOKEN_EXPIRY_STRING")
		t, err := time.Parse(time.RFC3339, TOKEN_EXPIRY_STRING)
		if err != nil {
			log.Fatalln("Invalid TOKEN_EXPIRY formatting", err)
		}
		TOKEN_EXPIRY = t
	}

	// GitHub Workflow 用
	if os.Getenv("HALT_AFTER_SUCCESS") == "true" {
		HALT_AFTER_SUCCESS = true
	}
}
