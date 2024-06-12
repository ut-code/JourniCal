package router

import (
	"github.com/ut-code/JourniCal/backend/diary"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Diary(g *echo.Group, db *gorm.DB) {
	g.GET("/", func(c echo.Context) error {
		return diary.GetAllDiaries(c, db)
	})

	g.GET("/:id", func(c echo.Context) error {
		return diary.GetDiaryByID(c, db)
	})

	g.POST("/", func(c echo.Context) error {
		return diary.CreateDiary(c, db)
	})

	g.PUT("/:id", func(c echo.Context) error {
		return diary.UpdateDiary(c, db)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		return diary.DeleteDiary(c, db)
	})
}
