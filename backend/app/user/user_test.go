package user_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)

	db, err := gorm.Open(sqlite.Open("./test.db"))
	db.AutoMigrate(&user.User{})
	helper.PanicIf(err)
	randomValue := "123456789"

	u, err := user.CreateUser(db, "USERNAME", "password", randomValue, randomValue)
	helper.PanicIf(err)
	u2, err := user.FindUserFromPassword(db, "USERNAME", "password")
	helper.PanicIf(err)
	assert.Equal(u2.Username, "USERNAME")
	assert.Equal(u2.ID, u.ID)

	_, err = user.FindUserFromPassword(db, "USERNAME", "password2")
	assert.NotNil(err)

	u4, err := user.FindUserFromSession(db, user.SessionUser{
		ID:       u.ID,
		Username: "USERNAME",
		Session:  u.Session,
	})
	helper.PanicIf(err)
	assert.Equal(u4.ID, u.ID)
	assert.Equal(u4.Username, u.Username)

	os.Remove("./test.db")
}
