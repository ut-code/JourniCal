package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"JourniCalBackend/helper"
	"JourniCalBackend/router"
)

func init() {
}

func main() {
	db := helper.Database

	// Doc: https://echo.labstack.com/
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
	}))

	e.Static("/", "./static")
	router.Root(e.Group(""))
	router.Api(e.Group("/api"))
	router.Auth(e.Group("/api/auth"))
	router.Calendar(e.Group("/api/calendar"))
	router.Diary(e.Group("/api/diaries"), db)

	// サーバの起動
	if err := e.Start(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
