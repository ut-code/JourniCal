package diary

// TODO: Please write some tests please

import (
	"errors"
	"time"

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
	EventID   string    `gorm:"unique" json:"eventId"`
}
type HTTPStatus = int

func Get(db *gorm.DB, diaryID uint, owner *user.User) (*Diary, error) {
	diary, err := GetUnchecked(db, diaryID)
	if err != nil {
		return nil, err
	}
	if diary.CreatorID != owner.ID {
		return nil, errors.New("this diary is not yours")
	}
	return diary, nil
}

func GetAll(db *gorm.DB, u *user.User) ([]Diary, error) {
	diaries, err := GetAllUnchecked(db, u.ID)
	if err != nil {
		return nil, errors.New("database error")
	}
	return diaries, nil
}

func GetByEvent(db *gorm.DB, eventID string, owner *user.User) (*Diary, error) {
	diary, err := GetByEventUnchecked(db, eventID)
	if err != nil {
		return nil, errors.New("diary not found")
	}
	if diary.CreatorID != owner.ID {
		return nil, errors.New("this diary is not yours")
	}
	return diary, nil
}

func Create(db *gorm.DB, diary *Diary, creator *user.User) error {
	diary.CreatorID = creator.ID
	err := CreateUnchecked(db, diary)
	if err != nil {
		return err
	}
	return nil
}

func Update(db *gorm.DB, id uint, newDiary *Diary, requester *user.User) (*Diary, error) {
	diary, err := Get(db, uint(id), requester)
	if err != nil {
		return nil, errors.New("not found")
	}
	// this branch should not be triggered but just in case
	if diary.CreatorID != requester.ID {
		return nil, errors.New("you don't own this diary")
	}
	err = UpdateUnchecked(db, id, newDiary)
	if err != nil {
		return nil, errors.New("failed to update diary")
	}
	return newDiary, nil
}

// UNSAFE: this doesn't verify the owner.
func GetUnchecked(db *gorm.DB, id uint) (*Diary, error) {
	diary := &Diary{}
	if err := db.First(diary, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return diary, nil
}

// UNSAFE: the username is not validated here.
func GetAllUnchecked(db *gorm.DB, creatorId uint) ([]Diary, error) {
	diaries := []Diary{}
	if err := db.Where("creator_id = ?", creatorId).Find(&diaries).Error; err != nil {
		return nil, errors.New("database error: failed to get diaries of a user")
	}
	return diaries, nil
}

func GetByEventUnchecked(db *gorm.DB, eventID string) (*Diary, error) {
	diary := &Diary{}
	if err := db.Where("event_id = ?", eventID).First(diary).Error; err != nil {
		return nil, errors.New("diary not found")
	}
	return diary, nil
}

func CreateUnchecked(db *gorm.DB, diary *Diary) error {
	if err := db.Create(diary).Error; err != nil {
		return errors.New("database error: failed to create diary")
	}
	return nil
}

func UpdateUnchecked(db *gorm.DB, id uint, newDiary *Diary) error {
	if id != newDiary.ID {
		return errors.New("unmatched ID")
	}
	diary := Diary{}
	if err := db.First(&diary, id).Error; err != nil {
		return err
	}
	if err := db.Save(newDiary).Error; err != nil {
		return err
	}
	return nil
}

func Delete(db *gorm.DB, id uint, owner *user.User) error {
	ds := []Diary{}
	if db.Where(`id = ?`, id).Where(`creator_id = ?`, owner.ID).Find(&ds); len(ds) == 0 {
		return errors.New("you don't own a diary that match the cond")
	}
	return db.Model(&Diary{}).Where(`id = ?`, id).Where(`creator_id = ?`, owner.ID).Delete(&Diary{}).Error
}
