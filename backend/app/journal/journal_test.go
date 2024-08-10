package journal_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ut-code/JourniCal/backend/app/journal"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"))
	helper.PanicOn(err)
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&journal.Journal{})
}

func TestBasicFunctionality(t *testing.T) {
	rand := "RANDOM_VALUE"
	u, err := user.Create(db, "username", "password", rand, rand, nil)
	helper.PanicOn(err)
	unauthorizedUser, err := user.Create(db, "another username", "another password", rand, "random string", nil)
	helper.PanicOn(err)
	assert := assert.New(t)

	d := &journal.Journal{
		Date:    time.Now(),
		Title:   "Hello, World!",
		Content: "Lorem Ipsum",
		EventID: "311c6584-812c-7f3c-182b-2880ce31af21",
	}
	err = journal.Create(db, d, u)
	assert.Nil(err)
	assert.Equal(u.ID, d.CreatorID)

	d2 := &journal.Journal{
		Date:    time.Now(),
		Title:   "Good morning, world",
		Content: "Consectetur adipiscing elit.",
		EventID: "2f2ee8d2-fae0-7dbf-c04a-f3a0999029d1",
	}
	err = journal.Create(db, d2, u)
	assert.Nil(err)
	assert.Equal(d.CreatorID, u.ID)

	dd, err := journal.Get(db, d.ID, u)
	helper.PanicOn(err)
	assert.Equal(d.ID, dd.ID, "d should equal dd")
	assert.Equal(d.Title, dd.Title, "d should equal dd")
	assert.Equal(d.Content, dd.Content, "d should equal dd")
	assert.Equal(d.EventID, dd.EventID, "d should equal dd")

	journals, err := journal.GetAll(db, u)
	helper.PanicOn(err)
	assert.Equal(2, len(journals), "len(journals) should equal 2")
	helper.PanicIf(len(journals) < 1)
	assert.Equal(d.ID, journals[0].ID, "journals[0] should equal d")
	assert.Equal(d.Content, journals[0].Content, "journals[0] should equal d")
	assert.Equal(d2.ID, journals[1].ID, "journals[1] should equal d2")
	assert.Equal(d2.Content, journals[1].Content, "journals[1] should equal d2")

	dd, err = journal.GetByEvent(db, d.EventID, u)
	helper.PanicOn(err)
	assert.Equal(d.ID, dd.ID, "d should equal dd")
	assert.Equal(d.Title, dd.Title, "d should equal dd")
	assert.Equal(d.Content, dd.Content, "d should equal dd")
	assert.Equal(d.EventID, dd.EventID, "d should equal dd")

	dd, err = journal.Get(db, d.ID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to GET")
	assert.Nil(dd, "unauthorized user should not be able to GET")

	journals, err = journal.GetAll(db, unauthorizedUser)
	assert.Nil(err)
	assert.Equal(len(journals), 0, "unauthorized user should not be able to get any")

	dd, err = journal.GetByEvent(db, d.EventID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to GET by event")
	assert.Nil(dd, "unauthorized user should not be able to GET by event")

	err = journal.Delete(db, d.ID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to delete")
	journals, err = journal.GetAll(db, u)
	assert.Nil(err)
	assert.Equal(2, len(journals), "unauthorized user should not be able to delete")

	err = journal.Delete(db, d.ID, u)
	assert.Nil(err)
	journals, err = journal.GetAll(db, u)
	assert.Nil(err)
	assert.Equal(1, len(journals), "test deletion")

	d, err = journal.Get(db, d.ID, u)
	assert.Error(err, "journal should've been deleted")
	assert.Nil(d)
}
