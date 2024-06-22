package user_test

import (
	"testing"

	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/test/assertion"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	assert := assertion.New(t)

	db, err := gorm.Open(sqlite.Open("./test.db"))
	assert.PanicOn(err)
	randomValue := "123456789"

	u, err := user.CreateUser(db, "USERNAME", "password", randomValue, randomValue)
	assert.Nil(err)
	u2, err := user.FindUserFromPassword(db, "USERNAME", "password")
	assert.Nil(err)
	assert.Eq(u2.Name, "USERNAME")
	assert.Eq(u2.ID, u.ID)

	u3, err := user.FindUserFromSession(db, user.SessionUser{
		ID:      u.ID,
		Name:    "USERNAME",
		Session: u.Session,
	})
	assert.Nil(err)
	assert.Eq(u3.ID, u.ID)
	assert.Eq(u3.Name, u.Name)
}
