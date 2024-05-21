package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"JourniCalBackend/helpers"
	"JourniCalBackend/routes"
)

func main() {
	// I chose to use Echo.
	// Doc: https://echo.labstack.com/
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
	}))

	// static (directory-based) serving
	e.Static("/", "./static")

	db, err := helpers.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	api := e.Group("/api")

	routes.RegisterDiaryRoutes(api, db)

	// サーバの起動
	if err := e.Start(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
