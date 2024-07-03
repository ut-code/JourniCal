package env

import (
	"os"

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

	// GitHub Workflow 用
	if os.Getenv("HALT_AFTER_SUCCESS") == "true" {
		HALT_AFTER_SUCCESS = true
	}
}
