# CI Test

## NOTIFY

currently not working.

waiting for google oauth2 long-term session key to be 
stored somewhere other than in-memory cache and reused.

## How to run

1. setup the project
2. copy credentials.json and .env to ./test/
3. use 
`curl -b token.cookie "http://localhost:3000/auth/code?state=state-token&code=YOUR_AUTH_CODE"`
to obtain token cookie
4. copy the token.cookie to ./test/
5. run `go test ./test/`
