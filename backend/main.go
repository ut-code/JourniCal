package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
)

// HTTPServer is top-level because it is an interface between client and has to be able to run every function.
func main() {
	// Doc: https://echo.labstack.com/
	e := echo.New()
	ctx = context.Background()
	cfg = ReadCredentials()

	// ミドルウェアを設定
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", getRoot)
	api := e.Group("/api")
	api.POST("/ping", postApiPing)

	e.GET("/auth/new", getAuthNew)
	e.GET("/auth/code", getAuthCode)
	e.GET("/auth/check", getAuthCheck)

	// listen + serve
	log.Fatal(e.Start(":3000"))
}
