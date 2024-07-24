package app

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ut-code/JourniCal/backend/app/database"
	"github.com/ut-code/JourniCal/backend/app/diary"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/router"
	"github.com/ut-code/JourniCal/backend/app/user"
)

var e *echo.Echo

func init() {
	db := db.InitDB(
		&diary.Diary{},
		&user.User{},
	)

	// Doc: https://echo.labstack.com/
	e = echo.New()
	// ミドルウェアを設定
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	mustLogin := user.LoginMiddleware(db)

	if options.ENABLE_CORS {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{options.CORS_ORIGIN},
			AllowCredentials: true,
		}))
	}
	if options.ECHO_SERVES_FRONTEND_TOO {
		e.Static("/", "./static")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello from Echo!")
	})

	router.Auth(e.Group("/auth"), db)
	router.User(e.Group("/api/user", mustLogin), db)
	router.Calendar(e.Group("/api/calendar", mustLogin), db)
	router.Diary(e.Group("/api/diaries", mustLogin), db)

	// GitHub CI 用
	if options.HALT_AFTER_SUCCESS {
		go func() {
			time.Sleep(15 * time.Second)
			os.Exit(0)
		}()
	}
}

func Serve(port uint) {
	// サーバの起動
	if err := e.Start(":" + fmt.Sprint(port)); err != nil {
		fmt.Println(err.Error())
	}
}
