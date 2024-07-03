package app

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"

	"github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/database"
	"github.com/ut-code/JourniCal/backend/app/diary"
	"github.com/ut-code/JourniCal/backend/app/env"
	"github.com/ut-code/JourniCal/backend/app/router"
	"github.com/ut-code/JourniCal/backend/app/user"
)

var e *echo.Echo
var conf *oauth2.Config
var AuthURL string

func init() {
	db := db.InitDB(
		&diary.Diary{},
		&user.User{},
	)
	if !env.NO_CREDENTIALS_REQUIRED {
		conf = calendar.ReadCredentials()
	}
	AuthURL = conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// Doc: https://echo.labstack.com/
	e = echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if cors_origin := os.Getenv("CORS_ORIGIN"); cors_origin != "" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{cors_origin},
			AllowCredentials: true,
		}))
	}
	if os.Getenv("ECHO_SERVES_FRONTEND_TOO") == "true" {
		e.Static("/", "./static")
	}
	router.Root(e.Group(""), db, AuthURL, conf)
	router.Api(e.Group("/api"))
	router.Auth(e.Group("/auth"), db, AuthURL, conf)
	router.User(e.Group("/api/user"), db)
	router.Calendar(e.Group("/api/calendar"), db, AuthURL, conf)
	router.Diary(e.Group("/api/diaries"), db)

	// GitHub CI 用
	if os.Getenv("HALT_AFTER_SUCCESS") == "true" {
		go func() {
			time.Sleep(15 * time.Second)
			os.Exit(0)
		}()
	}
}

func Serve(port int) {
	// サーバの起動
	if err := e.Start(":" + fmt.Sprint(port)); err != nil {
		fmt.Println(err.Error())
	}
}
