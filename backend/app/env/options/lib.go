package options

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// NOTE: this can also be written in .env, because godotenv is loaded twice.
// ALLOWS: multiple values separated by space.
// supported values: "secret" or "localtest". they will load ".env.secret" and "env/localtest.env" respectively.
var ENV_FILE string

var TOKEN_SOURCE TokenSource             // "db", "file", "env", or "none". defaults to "db" if not set.
var CREDENTIALS_SOURCE CredentialsSource // "file", "env" or "none". defaults to "file" if not set. if "none", auth-related things cannot be done.

var STATIC_USER = false              // whether to use static user for everything
var STATIC_TOKEN = false             // will be set to true if TOKEN_SOURCE != db. requires STATIC_USER to be set true.
var HALT_AFTER_SUCCESS = false       // GitHub Workflow 用
var DEV_ROUTES = false               // whether to  enable routes meant for dev-use only.
var ECHO_SERVES_FRONTEND_TOO = false // whether echo serves ./static/ as well as backend.
var PREFILL_JOURNAL = false          // whether to prefill journal on startup
var IN_MEMORY_DB = false             // whether to use sqlite's in-memory db.

var CORS_ORIGIN string  // optional
var ENABLE_CORS = false // will be set true if CORS_ORIGIN != "". cannot directly modify from env.

type TokenSource int

const (
	TokenSourceDB TokenSource = iota // also set if it's "none"
	TokenSourceEnv
	TokenSourceFile
)

type CredentialsSource int

const (
	CredentialsSourceFile CredentialsSource = iota
	CredentailsSourceEnv
	CredentialsSourceNone
)

func init() {
	// load .env once for ENV_FILE
	godotenv.Load(".env")

	var envfile []string
	for _, filename := range strings.Split(os.Getenv("ENV_FILE"), " ") {
		var appending string
		switch strings.ToLower(filename) {
		case "":
			continue // not breaking in case of multi space in between such as ENV_FILE="secret    test"
		case ".env":
			log.Fatalln("You don't need to specify .env because it is already loaded.")
		case "secret", ".env.secret":
			appending = ".env.secret"
		case "test", "localtest", "localtest.env":
			appending = "env/localtest.env"
		case "dev", "dev.env":
			appending = "env/dev.env"
		default:
			log.Fatalln("ERROR: assertion failed in app/env/options: unknown ENV_FILE string:", filename)
		}
		envfile = append(envfile, appending)
	}
	if len(envfile) >= 1 {
		godotenv.Load(envfile...)
	}

	STATIC_USER = boolean("STATIC_USER")
	HALT_AFTER_SUCCESS = boolean("HALT_AFTER_SUCCESS")
	DEV_ROUTES = boolean("DEV_ROUTES")
	ECHO_SERVES_FRONTEND_TOO = boolean("ECHO_SERVES_FRONTEND_TOO")
	PREFILL_JOURNAL = boolean("PREFILL_JOURNAL")
	IN_MEMORY_DB = boolean("IN_MEMORY_DB")

	if corsOrigin := env("CORS_ORIGIN"); corsOrigin != "" {
		ENABLE_CORS = true
		CORS_ORIGIN = corsOrigin
	}

	switch src := env("TOKEN_SOURCE"); strings.ToLower(src) {
	case "db", "database", "", "none":
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

	switch src := env("CREDENTIALS_SOURCE"); strings.ToLower(src) {
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

	if STATIC_USER && !STATIC_TOKEN {
		log.Fatalln("validation failed: TOKEN_SOURCE must be set to either file or env when using STATIC_USER")
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
