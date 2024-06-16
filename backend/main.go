package main

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ut-code/JourniCal/backend/diary"
	"github.com/ut-code/JourniCal/backend/helper/database"
	"github.com/ut-code/JourniCal/backend/router"
)

func init() {
}

func main() {
	diaryDB := db.InitDB(&diary.Diary{})
	// Doc: https://echo.labstack.com/
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if cors_origin := os.Getenv("CORS_ORIGIN"); cors_origin != "" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{cors_origin},
		}))
	}
	if os.Getenv("ECHO_SERVES_FRONTEND_TOO") == "true" {
		e.Static("/", "./static")
	}
	router.Root(e.Group(""))
	router.Api(e.Group("/api"))
	router.Auth(e.Group("/api/auth"))
	router.Calendar(e.Group("/api/calendar"))
	router.Diary(e.Group("/api/diaries"), diaryDB)

	// GitHub CI 用
	if os.Getenv("HALT_AFTER_SUCCESS") == "true" {
		go func() {
			time.Sleep(15 * time.Second)
			os.Exit(0)
		}()
	}

	// サーバの起動
	if err := e.Start(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
