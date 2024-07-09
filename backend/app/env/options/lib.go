package options

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var TOKEN_SOURCE TokenSource             // "db", "file" or "env". defaults to "db" if not set.
var CREDENTIALS_SOURCE CredentialsSource // "file", "env" or "none". defaults to "file" if not set. if "none", auth-related things cannot be done.

var STATIC_USER = false              // whether to use static user for everything
var STATIC_TOKEN = false             // will be set to true if TOKEN_SOURCE != db. requires STATIC_USER to be set true.
var ECHO_SERVES_FRONTEND_TOO = false // whether echo serves ./static/ as well as backend.
var HALT_AFTER_SUCCESS = false       // GitHub Workflow 用
var ENABLE_CORS = false              // will be set true if CORS_ORIGIN != "". cannot directly modify from env.

var CORS_ORIGIN string // optional

type TokenSource int
type CredentialsSource int

const (
	TokenSourceDB TokenSource = iota
	TokenSourceEnv
	TokenSourceFile
)

const (
	CredentialsSourceFile = iota
	CredentailsSourceEnv
	CredentialsSourceNone
)

func init() {
	godotenv.Load()

	STATIC_USER = boolean("STATIC_USER")
	ECHO_SERVES_FRONTEND_TOO = boolean("ECHO_SERVES_FRONTEND_TOO")
	HALT_AFTER_SUCCESS = boolean("HALT_AFTER_SUCCESS")

	if corsOrigin := env("CORS_ORIGIN"); corsOrigin != "" {
		ENABLE_CORS = true
		CORS_ORIGIN = corsOrigin
	}

	switch src := env("TOKEN_SOURCE"); src {
	case "db", "database", "":
		TOKEN_SOURCE = TokenSourceDB
	case "env", "environment":
		TOKEN_SOURCE = TokenSourceEnv
		STATIC_TOKEN = true
	case "file":
		TOKEN_SOURCE = TokenSourceFile
		STATIC_TOKEN = true
	default:
		log.Fatalln("Failed assertion in TOKEN_SOURCE.\n  - Must be one of: db, env, file\n  - Got: " + src)
	}

	switch src := env("CREDENTIALS_SOURCE"); src {
	case "file", "":
		CREDENTIALS_SOURCE = CredentialsSourceFile
	case "env", "environment":
		CREDENTIALS_SOURCE = CredentailsSourceEnv
	case "none", "nil":
		CREDENTIALS_SOURCE = CredentialsSourceNone
	default:
		log.Fatalln("Failed assertion in CREDENTIALS_SOURCE.\n  - Must be one of: file, env\n  - Got: " + src)
	}

	// validation

	if STATIC_TOKEN && !STATIC_USER {
		log.Fatalln("validation failed; STATIC_USER must be set true when using STATIC_TOKEN")
	}
}

// returns $name.
// doesn't panic if $name == "" or $name is not set (returns "").
func env(name string) string {
	return os.Getenv(name)
}

// this returns true if $name is "true".
// returns false otherwise.
func boolean(name string) bool {
	env := os.Getenv(name)
	return env == "true"
}
