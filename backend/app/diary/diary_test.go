package diary_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ut-code/JourniCal/backend/app/diary"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	_ "github.com/ut-code/JourniCal/backend/pkg/tests/run-test-at-root"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"))
	helper.PanicOn(err)
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&diary.Diary{})
}

func TestBasicFunctionality(t *testing.T) {
	rand := "RANDOM_VALUE"
	u, err := user.Create(db, "username", "password", rand, rand, nil)
	helper.PanicOn(err)
	unauthorizedUser, err := user.Create(db, "another username", "another password", rand, "random string", nil)
	helper.PanicOn(err)
	assert := assert.New(t)

	d := &diary.Diary{
		Date:    time.Now(),
		Title:   "Hello, World!",
		Content: "Lorem Ipsum",
		EventID: "311c6584-812c-7f3c-182b-2880ce31af21",
	}
	err = diary.Create(db, d, u)
	assert.Nil(err)
	assert.Equal(u.ID, d.CreatorID)

	d2 := &diary.Diary{
		Date:    time.Now(),
		Title:   "Good morning, world",
		Content: "Consectetur adipiscing elit.",
		EventID: "2f2ee8d2-fae0-7dbf-c04a-f3a0999029d1",
	}
	err = diary.Create(db, d2, u)
	assert.Nil(err)
	assert.Equal(d.CreatorID, u.ID)

	dd, err := diary.Get(db, d.ID, u)
	helper.PanicOn(err)
	assert.Equal(d.ID, dd.ID, "d should equal dd")
	assert.Equal(d.Title, dd.Title, "d should equal dd")
	assert.Equal(d.Content, dd.Content, "d should equal dd")
	assert.Equal(d.EventID, dd.EventID, "d should equal dd")

	diaries, err := diary.GetAll(db, u)
	helper.PanicOn(err)
	assert.Equal(2, len(diaries), "len(diaries) should equal 2")
	helper.PanicIf(len(diaries) < 1)
	assert.Equal(d.ID, diaries[0].ID, "diaries[0] should equal d")
	assert.Equal(d.Content, diaries[0].Content, "diaries[0] should equal d")
	assert.Equal(d2.ID, diaries[1].ID, "diaries[1] should equal d2")
	assert.Equal(d2.Content, diaries[1].Content, "diaries[1] should equal d2")

	dd, err = diary.GetByEvent(db, d.EventID, u)
	helper.PanicOn(err)
	assert.Equal(d.ID, dd.ID, "d should equal dd")
	assert.Equal(d.Title, dd.Title, "d should equal dd")
	assert.Equal(d.Content, dd.Content, "d should equal dd")
	assert.Equal(d.EventID, dd.EventID, "d should equal dd")

	dd, err = diary.Get(db, d.ID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to GET")
	assert.Nil(dd, "unauthorized user should not be able to GET")

	diaries, err = diary.GetAll(db, unauthorizedUser)
	assert.Nil(err)
	assert.Equal(len(diaries), 0, "unauthorized user should not be able to get any")

	dd, err = diary.GetByEvent(db, d.EventID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to GET by event")
	assert.Nil(dd, "unauthorized user should not be able to GET by event")

	err = diary.Delete(db, d.ID, unauthorizedUser)
	assert.Error(err, "unauthorized user should not be able to delete")
	diaries, err = diary.GetAll(db, u)
	assert.Nil(err)
	assert.Equal(2, len(diaries), "unauthorized user should not be able to delete")

	err = diary.Delete(db, d.ID, u)
	assert.Nil(err)
	diaries, err = diary.GetAll(db, u)
	assert.Nil(err)
	assert.Equal(1, len(diaries), "test deletion")

	d, err = diary.Get(db, d.ID, u)
	assert.Error(err, "diary should've been deleted")
	assert.Nil(d)
}
