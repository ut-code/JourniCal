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
	ID        uint      `json:"id"`
	CreatorID uint      `json:"creatorId"`
	Date      time.Time `json:"date"` // Date of what?
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}
type HTTPStatus = int

func GetAllOfUser(c echo.Context, db *gorm.DB) (HTTPStatus, []Diary, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Authentication error")
	}
	diaries, err := GetAllUnchecked(db, u.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("database error")
	}
	return 200, diaries, nil
}

// UNSAFE: the username is not validated here.
func GetAllUnchecked(db *gorm.DB, creatorId uint) ([]Diary, error) {
	diaries := []Diary{}
	if err := db.Where("creator_id = ?", creatorId).Find(&diaries).Error; err != nil {
		return nil, errors.New("Database error: failed to get diaries of a user")
	}
	return diaries, nil
}

func GetUnchecked(db *gorm.DB, id uint) (*Diary, error) {
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return diary, nil
}

func Get(c echo.Context, db *gorm.DB) (HTTPStatus, *Diary, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Authentication error")
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return http.StatusBadRequest, nil, errors.New("Invalid id format")
	}
	diary, err := GetUnchecked(db, uint(id))
	if diary.CreatorID != u.ID {
		return http.StatusNotFound, nil, errors.New("This diary is not yours")
	}
	return http.StatusOK, diary, nil
}

func Create(c echo.Context, db *gorm.DB) (HTTPStatus, *Diary, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Invalid user. consider creating new one.")
	}
	diary := new(Diary)
	// UNTESTED: I'm not sure how this works or how this should be done. Test this function.
	if err := c.Bind(diary); err != nil {
		return http.StatusBadRequest, nil, errors.New("Failed to bind diary data")
	}
	diary.CreatorID = u.ID
	err = CreateUnchecked(db, diary)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusCreated, diary, nil
}

func CreateUnchecked(db *gorm.DB, diary *Diary) error {
	if err := db.Create(diary).Error; err != nil {
		return errors.New("Database error: Failed to create diary")
	}
	return nil
}

func Update(c echo.Context, db *gorm.DB) (HTTPStatus, *Diary, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Invalid user. create new one")
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return http.StatusBadRequest, nil, errors.New("Invalid formating of id or negative value")
	}
	diary, err := GetUnchecked(db, uint(id))
	if err != nil {
		return http.StatusNotFound, nil, errors.New("not found")
	}
	if diary.CreatorID != u.ID {
		return http.StatusUnauthorized, nil, errors.New("You don't own this diary")
	}

	var newDiary Diary
	if err := c.Bind(&newDiary); err != nil {
		return http.StatusBadRequest, nil, errors.New("Failde to bind diary")
	}
	err = UpdateUnchecked(db, uint(id), newDiary)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("Failed to update diary")
	}
	return http.StatusOK, diary, nil
}

func UpdateUnchecked(db *gorm.DB, id uint, newDiary Diary) error {
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

func Delete(c echo.Context, db *gorm.DB) (HTTPStatus, error) {
	u, err := user.FromEchoContext(db, c)
	if err != nil {
		return http.StatusBadRequest, errors.New("Authentication error: user not found")
	}
	intid, err := strconv.Atoi(c.Param("id"))
	if err != nil || intid < 0 {
		return http.StatusBadRequest, errors.New("Invalid diary ID")
	}
	id := uint(intid)
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return http.StatusNotFound, errors.New("Diary not found")
	}
	if diary.CreatorID != u.ID {
		return http.StatusUnauthorized, errors.New("You don't own this diary")
	}
	if err := DeleteUnchecked(db, id); err != nil {
		return http.StatusInternalServerError, errors.New("Failed to delete diary: " + err.Error())
	}
	return http.StatusNoContent, nil
}

func DeleteUnchecked(db *gorm.DB, id uint) error {
	return db.Delete(&Diary{}, `id = ?`, id).Error
}
