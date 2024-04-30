package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type api_root_res struct {
	// structs must start with capital letters to allow JSON stringify (called JSON.Marshal) to access the property
	// in Golang, all fields that start with lower-case letters are considered private.
	Hello string
	Query url.Values
}

func main() {
	// CalendarSample()
	HTTPServerSample()
}

// HTTPServer is top-level because it is an interface between client and has to be able to run every function.
func HTTPServerSample() {
	// I chose to use Echo.
	// Doc: https://echo.labstack.com/
	e := echo.New()

	// ミドルウェアを設定
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	// static (directory-based) serving
	e.Static("/", "./static")

	// file-based serving
	e.File("/", "index.html")
	// or more verbosely, (replace POST with GET for the same result)
	e.POST("/", func(c echo.Context) error {
		err := c.String(http.StatusOK, "Alternative Text (because reading file in Go is a pain)")
		// normal error handling example for beginners.
		if err != nil {
			// handle err here
			return err
		}
		return nil // returning nil as error means the operation was successful.

		// OR, you can also just do return c.String(...) to return error directly.
	})

	// this will create a sub-route under /api/
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
	api.GET("/diaries", func(c echo.Context) error {
		diaries := GetDiary()
		c.JSON(http.StatusOK, diaries)
		return nil
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

	// write any code here

	// listen + serve
	err := e.Start(":3000")
	fmt.Println(err.Error())
}
