package env

import (
	"os"

	"github.com/joho/godotenv"
)

var USE_TOKEN_JSON = false
var STATIC_USER = false
var NO_CREDENTIALS_REQUIRED = false

// WIP

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
}
