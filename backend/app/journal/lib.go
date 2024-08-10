package journal

// TODO: Please write some tests please

import (
	"errors"
	"time"

	"github.com/ut-code/JourniCal/backend/app/user"
	"gorm.io/gorm"
)

type Journal struct {
	gorm.Model
	ID        uint      `json:"id"`
	CreatorID uint      `json:"creatorId"`
	Date      time.Time `json:"date"` // Date of what?
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	EventID   string    `gorm:"unique" json:"eventId"`
}
type HTTPStatus = int

func Get(db *gorm.DB, journalID uint, owner *user.User) (*Journal, error) {
	journal, err := GetUnchecked(db, journalID)
	if err != nil {
		return nil, err
	}
	if journal.CreatorID != owner.ID {
		return nil, errors.New("this journal is not yours")
	}
	return journal, nil
}

func GetAll(db *gorm.DB, u *user.User) ([]Journal, error) {
	journals, err := GetAllUnchecked(db, u.ID)
	if err != nil {
		return nil, errors.New("database error")
	}
	return journals, nil
}

func GetByEvent(db *gorm.DB, eventID string, owner *user.User) (*Journal, error) {
	journal, err := GetByEventUnchecked(db, eventID)
	if err != nil {
		return nil, errors.New("journal not found")
	}
	if journal.CreatorID != owner.ID {
		return nil, errors.New("this journal is not yours")
	}
	return journal, nil
}

// updates journal.ID and journal.CreatorID
func Create(db *gorm.DB, journal *Journal, creator *user.User) error {
	journal.CreatorID = creator.ID
	err := CreateUnchecked(db, journal)
	if err != nil {
		return err
	}
	return nil
}

func Update(db *gorm.DB, id uint, newJournal *Journal, requester *user.User) (*Journal, error) {
	journal, err := Get(db, uint(id), requester)
	if err != nil {
		return nil, errors.New("not found")
	}
	// this branch should not be triggered but just in case
	if journal.CreatorID != requester.ID {
		return nil, errors.New("you don't own this journal")
	}
	err = UpdateUnchecked(db, id, newJournal)
	if err != nil {
		return nil, errors.New("failed to update journal")
	}
	return newJournal, nil
}

// UNSAFE: this doesn't verify the owner.
func GetUnchecked(db *gorm.DB, id uint) (*Journal, error) {
	journal := &Journal{}
	if err := db.First(journal, id).Error; err != nil {
		return nil, errors.New("not found")
	}
	return journal, nil
}

// UNSAFE: the username is not validated here.
func GetAllUnchecked(db *gorm.DB, creatorId uint) ([]Journal, error) {
	journals := []Journal{}
	if err := db.Where("creator_id = ?", creatorId).Find(&journals).Error; err != nil {
		return nil, errors.New("database error: failed to get journals of a user")
	}
	return journals, nil
}

func GetByEventUnchecked(db *gorm.DB, eventID string) (*Journal, error) {
	journal := &Journal{}
	if err := db.Where("event_id = ?", eventID).First(journal).Error; err != nil {
		return nil, errors.New("journal not found")
	}
	return journal, nil
}

// updates journal.ID
func CreateUnchecked(db *gorm.DB, journal *Journal) error {
	if err := db.Create(journal).Error; err != nil {
		return errors.New("database error: failed to create journal")
	}
	return nil
}

func UpdateUnchecked(db *gorm.DB, id uint, newJournal *Journal) error {
	if id != newJournal.ID {
		return errors.New("unmatched ID")
	}
	journal := Journal{}
	if err := db.First(&journal, id).Error; err != nil {
		return err
	}
	if err := db.Save(newJournal).Error; err != nil {
		return err
	}
	return nil
}

func Delete(db *gorm.DB, id uint, owner *user.User) error {
	js := []Journal{}
	if db.Where(`id = ?`, id).Where(`creator_id = ?`, owner.ID).Find(&js); len(js) == 0 {
		return errors.New("you don't own a journal that match the cond")
	}
	return db.Model(&Journal{}).Where(`id = ?`, id).Where(`creator_id = ?`, owner.ID).Delete(&Journal{}).Error
}
