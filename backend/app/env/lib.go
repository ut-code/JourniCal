package env

import "github.com/joho/godotenv"

var USE_TOKEN_JSON = false

// WIP

func init() {
	godotenv.Load()
}
