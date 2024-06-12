package diary

import (
	"github.com/ut-code/JourniCal/backend/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Diary = types.Diary

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

func UpdateDiary(c echo.Context, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid diary ID"})
	}
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Diary not found"})
	}
	if err := c.Bind(diary); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to bind diary data"})
	}
	if err := db.Save(diary).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update diary"})
	}
	return c.JSON(http.StatusOK, diary)
}

func DeleteDiary(c echo.Context, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid diary ID"})
	}
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Diary not found"})
	}
	if err := db.Delete(diary).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete diary"})
	}
	return c.NoContent(http.StatusNoContent)
}
