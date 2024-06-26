package diary

// TODO: Please write some tests please

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ut-code/JourniCal/backend/app/user"
	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	Creator   user.User `json:"creator" gorm:"foreignKey:CreatorID;references:ID"`
	CreatorID int
	Date      time.Time `json:"date"` // Date of what?
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func GetAllDiariesOfUser(c echo.Context, db *gorm.DB) error {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Authentication error"})
	}
	diaries, err := UncheckedGetAllDiariesOfUsername(db, u.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}
	return c.JSON(http.StatusOK, diaries)
}

// UNSAFE: the username is not validated here.
func UncheckedGetAllDiariesOfUsername(db *gorm.DB, creator string) ([]Diary, error) {
	diaries := []Diary{}
	if err := db.Where("creator = ?", creator).Find(&diaries).Error; err != nil {
		return nil, errors.New("Database error: failed to get diaries of a user")
	}
	return diaries, nil
}

func UncheckedGetDiaryByID(db *gorm.DB, id int) (*Diary, error) {
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return diary, nil
}
func GetDiaryByID(c echo.Context, db *gorm.DB) error {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Authentication error"})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id format"})
	}
	diary, err := UncheckedGetDiaryByID(db, id)
	if diary.CreatorID != u.ID {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "This diary is not yours"})
	}
	return c.JSON(http.StatusOK, diary)
}

func CreateDiary(c echo.Context, db *gorm.DB) error {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user. consider creating new one."})
	}
	diary := new(Diary)
	// UNTESTED: I'm not sure how this works or how this should be done. Test this function.
	diary.Creator = *u
	if err := c.Bind(diary); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to bind diary data"})
	}
	if err := db.Create(diary).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create diary"})
	}
	return c.JSON(http.StatusCreated, diary)
}

func UpdateDiary(c echo.Context, db *gorm.DB) error {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user. create new one"})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid formating of id or negative value"})
	}
	var newDiary Diary
	if err := c.Bind(newDiary); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failde to bind diary"})
	}
	diary, err := UncheckedGetDiaryByID(db, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "not found"})
	}
	if diary.CreatorID != u.ID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "You don't own this diary"})
	}
	err = UncheckedUpdateDiary(db, uint(id), newDiary)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update diary"})
	}
	return c.JSON(http.StatusOK, diary)
}

func UncheckedUpdateDiary(db *gorm.DB, id uint, newDiary Diary) error {
	if id != newDiary.ID {
		return errors.New("unmatched ID")
	}
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return err
	}
	if err := db.Save(newDiary).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDiary(c echo.Context, db *gorm.DB) error {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Authentication error: user not found"})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid diary ID"})
	}
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Diary not found"})
	}
	if diary.CreatorID != u.ID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "You don't own this diary"})
	}
	if err := db.Delete(diary).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete diary"})
	}
	return c.NoContent(http.StatusNoContent)
}
