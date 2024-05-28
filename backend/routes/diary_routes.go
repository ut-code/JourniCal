package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"JourniCalBackend/helpers"
)

func RegisterDiaryRoutes(api *echo.Group, db *gorm.DB) {
	api.GET("/diaries", func(c echo.Context) error {
		return helpers.GetAllDiaries(c, db)
	})

	api.GET("/diaries/:id", func(c echo.Context) error {
		return helpers.GetDiaryByID(c, db)
	})

	api.POST("/diaries", func(c echo.Context) error {
		return helpers.CreateDiary(c, db)
	})

	api.PUT("/diaries/:id", func(c echo.Context) error {
		return helpers.UpdateDiary(c, db)
	})

	api.DELETE("/diaries/:id", func(c echo.Context) error {
		return helpers.DeleteDiary(c, db)
	})
}
