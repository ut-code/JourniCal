package router

import (
	"github.com/ut-code/JourniCal/backend/app/diary"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Diary(g *echo.Group, db *gorm.DB) {
	g.GET("/", func(c echo.Context) error {
		status, json, err := diary.GetAllDiariesOfUser(c, db)
		if err != nil {
			return c.JSON(status, echo.Map{"error": err.Error()})
		}
		return c.JSON(status, json)
	})

	g.GET("/:id", func(c echo.Context) error {
		status, diary, err := diary.GetDiaryByID(c, db)
		if err != nil {
			return c.JSON(status, echo.Map{"error": err.Error()})
		}
		return c.JSON(status, diary)
	})

	g.POST("/", func(c echo.Context) error {
		status, diary, err := diary.CreateDiary(c, db)
		if err != nil {
			return c.JSON(status, echo.Map{"error": err.Error()})
		}
		return c.JSON(status, diary)
	})

	g.PUT("/:id", func(c echo.Context) error {
		status, diary, err := diary.UpdateDiary(c, db)
		if err != nil {
			return c.JSON(status, echo.Map{"error": err.Error()})
		}
		return c.JSON(status, diary)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		status, err := diary.DeleteDiary(c, db)
		if err != nil {
			return c.JSON(status, echo.Map{"error": err.Error()})
		}
		return c.NoContent(status)
	})
}
