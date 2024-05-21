package helpers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

func GetAllDiaries(c echo.Context, db *gorm.DB) error {
	diaries := []Diary{}
	if err := db.Find(&diaries).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, diaries)
}

func GetDiaryByID(c echo.Context, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid diary ID"})
	}
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Diary not found"})
	}
	return c.JSON(http.StatusOK, diary)
}

func CreateDiary(c echo.Context, db *gorm.DB) error {
	diary := new(Diary)
	if err := c.Bind(diary); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to bind diary data"})
	}
	if err := db.Create(diary).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create diary"})
	}
	return c.JSON(http.StatusCreated, diary)
}
