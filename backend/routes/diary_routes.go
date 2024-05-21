package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"JourniCalBackend/helpers"
)

func RegisterDiaryRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/diaries", func(c echo.Context) error {
		return helpers.GetAllDiaries(c, db)
	})

	e.GET("/diaries/:id", func(c echo.Context) error {
		return helpers.GetDiaryByID(c, db)
	})

	e.POST("/diaries", func(c echo.Context) error {
		return helpers.CreateDiary(c, db)
	})
}
