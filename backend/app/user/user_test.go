package user_test

import (
	"os"
	"testing"

	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/test/assertion"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	os.Remove("./test.db")
	assert := assertion.New(t)

	db, err := gorm.Open(sqlite.Open("./test.db"))
	db.AutoMigrate(&user.User{})
	assert.PanicOn(err)
	randomValue := "123456789"

	u, err := user.CreateUser(db, "USERNAME", "password", randomValue, randomValue)
	assert.PanicOn(err)
	u2, err := user.FindUserFromPassword(db, "USERNAME", "password")
	assert.PanicOn(err)
	assert.Eq(u2.Username, "USERNAME")
	assert.Eq(u2.ID, u.ID)

	_, err = user.FindUserFromPassword(db, "USERNAME", "password2")
	assert.NotNil(err)

	u4, err := user.FindUserFromSession(db, user.SessionUser{
		ID:       u.ID,
		Username: "USERNAME",
		Session:  u.Session,
	})
	assert.PanicOn(err)
	assert.Eq(u4.ID, u.ID)
	assert.Eq(u4.Username, u.Username)
}
