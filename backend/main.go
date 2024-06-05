package main

import (
	"JourniCalBackend/helper"
	"fmt"

	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"JourniCalBackend/helpers"
	"JourniCalBackend/routes"
	"golang.org/x/oauth2"
	"log"
)

var cfg *oauth2.Config
var ctx context.Context
var authURL string
var tokenCache = helper.NewMap[string, oauth2.Token]()

func init() {
	ctx = context.Background()
	cfg = ReadCredentials()
	authURL = cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func main() {
	// Doc: https://echo.labstack.com/
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
	}))

	e.GET("/", getRoot)
	api := e.Group("/api")
	api.POST("/ping", postApiPing)

	// spec:
	// specify time as unix time.
	api.GET("/get-20-events-forward/:start_unix", getGet20EventsForward)

	e.GET("/auth/new", getAuthNew)
	e.GET("/auth/code", getAuthCode)
	e.GET("/auth/check", getAuthCheck)

	// listen + serve
	log.Fatal(e.Start(":3000"))
	// static (directory-based) serving
	e.Static("/", "./static")

	db, err := helpers.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	routes.RegisterDiaryRoutes(api, db)

	// サーバの起動
	if err := e.Start(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
