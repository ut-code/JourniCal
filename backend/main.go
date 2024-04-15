package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// this struct is a sample.
type api_root_res struct {
	// structs must start with capital letters to allow JSON stringify (called JSON.Marshal) to access the property
	// in Golang, all fields that start with lower-case letters are considered private.
	Hello string
	Query url.Values
}

func main() {
	// do this at top-level initialization in main (or, copy & paste from inside the function)
	ctx := context.Background()
	cfg := ReadCredentials()
	token, err := tokenFromFile("./token.json")

	if err != nil {
		// token isn't there, therefore ask for token
		authURL := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Println("go to the link below and paste the `code` prop in the terminal:")
		fmt.Println(authURL)
		var b []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			b = append(b, scanner.Bytes()...)
			if scanner.Text() != "" {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}

		code := string(b)
		var err error
		token, err = cfg.Exchange(ctx, code)
		ErrorLog(err, "Unable to retrieve token from web")

		b, err = json.Marshal(token)
		ErrorLog(err)
		writeFile("./token.json", b)
	}
	if !token.Valid() {
		// expired token
		fmt.Println("Expired token. refreshing...")
		cfg.Client(ctx, token)
		b, err := json.Marshal(token)
		ErrorLog(err)
		writeFile("./token.json", b)
	}

	CalendarSample(ctx, *cfg, token)
	url, _ := readFile("./sample.url")
	GDriveSample(ctx, *cfg, token, url)
	// HTTPServerSample()
}

// HTTPServer is top-level because it is an interface between client and has to be able to run every function.
func HTTPServerSample() {
	// I chose to use Echo.
	// Doc: https://echo.labstack.com/
	e := echo.New()

	// google auth stuff
	ctx := context.Background()
	cfg := ReadCredentials()

	// ミドルウェアを設定
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	readAuthCode := func(c echo.Context) authCode {
		code, err := c.Cookie("code")
		if err != nil || code.Value == "" {
			SendToAuth(c, cfg)
			return ""
		}
		fmt.Println(code)
		return code.Value
	}
	Use(readAuthCode)

	// static (directory-based) serving
	e.GET("/", func(c echo.Context) error {
		readAuthCode(c)
		c.File("./index.html")
		return nil
	})
	// this will create a sub-route under /api
	api := e.Group("/api")
	// for example, this will handle a GET request to /api (careful, it won't handle a request to /api/ )
	api.GET("", func(c echo.Context) error {
		q := c.QueryParams()
		fmt.Println("request to /api was made. Query is: ", q)
		// JSON method takes object of any class and Marshals their *public* properties into res body.
		err := c.JSON(http.StatusOK, api_root_res{
			Hello: "World!",
			Query: q,
		})
		return err
	})
	// and this will handle a POST request to /api/ping (pong!)
	// try this in console: $ curl -X POST http://localhost:3000/api/ping -d "Hello there!"
	api.POST("/ping", func(c echo.Context) error {
		// read request body to bytes then stringify it
		// if you want to receive as JSON, there is a better way to do it:
		// json_map := make(map[string]interface{})
		// err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		buf, err := io.ReadAll(c.Request().Body)
		ErrorLog(err)
		body := string(buf)
		fmt.Println(body)
		c.String(http.StatusOK, "pong!\n Your request body was: \n"+body+"\n")
		return nil
	})

	e.GET("/auth/code", func(c echo.Context) error {
		code := c.QueryParam("code")
		if code == "" {
			c.String(http.StatusBadRequest, "empty authorization code")
			return nil // anyone can freely send a request to /auth/code so it's not an actual error
		}
		if ck, err := c.Cookie("code"); err == nil && ck.Value == code {
			c.Redirect(http.StatusOK, "/")
			return nil
		}

		tok, err := cfg.Exchange(ctx, code)
		SaveToken(code, tok)
		ErrorLog(err, "Unable to retrieve token from web")

		const MaxAge = 12 * 30 * 24 * 60 * 60 // about 3 months.
		c.SetCookie(&http.Cookie{
			Path:     "/",
			Name:     "code",
			Value:    code,
			MaxAge:   MaxAge,
			HttpOnly: true, // reduces XSS risk via disallowing access from browser JS
		})
		c.Redirect(http.StatusFound, "/")
		return nil
	})

	e.GET("/auth/check", func(c echo.Context) error {
		code, err := c.Cookie("code")
		if err != nil {
			c.String(200, "You are not authenticated.")
		} else {
			c.String(200, "You are authenticated. the code is: "+code.Value)
		}
		return nil
	})
	// write any code here

	// listen + serve
	err := e.Start(":3000")
	fmt.Println(err.Error())
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
