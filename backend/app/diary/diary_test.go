package diary_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ut-code/JourniCal/backend/app/diary"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
)

var db *gorm.DB

func init() {
	var err error
	os.Remove("./test.db")
	db, err = gorm.Open(sqlite.Open("./test.db"))
	helper.PanicOn(err)
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&diary.Diary{})
}

func TestBasicFunctionality(t *testing.T) {
	rand := "RANDOM_VALUE"
	u, err := user.CreateUser(db, "username", "password", rand, rand)
	helper.PanicOn(err)
	assert := assert.New(t)

	d := &diary.Diary{
		CreatorID: u.ID,
		Date:      time.Now(),
		Title:     "Hello, World!",
		Content:   "Lorem Ipsum",
	}
	err = diary.CreateUnchecked(db, d)
	assert.Nil(err)

	d2, err := diary.GetUnchecked(db, d.ID)
	helper.PanicOn(err)
	assert.Equal(d.ID, d2.ID, "d should equal d2")
	assert.Equal(d.Content, d2.Content, "d should equal d2")

	d3s, err := diary.GetAllUnchecked(db, u.ID)
	helper.PanicOn(err)
	assert.Equal(len(d3s), 1, "len(d3s) should equal 1")
	helper.PanicIf(len(d3s) < 1)
	assert.Equal(d3s[0].ID, d.ID, "d3s[0] should equal d")
	assert.Equal(d3s[0].Content, d.Content, "d3s[0] should equal d")
}
