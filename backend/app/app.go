package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	db "github.com/ut-code/JourniCal/backend/app/database"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/journal"
	"github.com/ut-code/JourniCal/backend/app/router"
	"github.com/ut-code/JourniCal/backend/app/user"
	echohandler "github.com/ut-code/JourniCal/backend/pkg/echo-handler"
)

var e *echo.Echo

func init() {
	db := db.InitDB(
		&journal.Journal{},
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
	if options.PREFILL_JOURNAL {
		journal.Prefill(db)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello from Echo!")
	})

	router.Auth(e.Group("/auth"), db)
	router.User(e.Group("/api/user", mustLogin), db)
	router.Calendar(e.Group("/api/calendar", mustLogin), db)
	router.Journal(e.Group("/api/journals", mustLogin), db)

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
	err := echohandler.Start(e, uint16(port))
	if err != nil {
		log.Fatalln(err)
	}
}
