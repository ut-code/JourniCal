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
	helper.PanicOn(err)
	randomValue := "123456789"

	u, err := user.CreateUser(db, "USERNAME", "password", randomValue, randomValue, nil)
	helper.PanicOn(err)

	_, err = user.CreateUser(db, "USERNAME", "different_password", randomValue, randomValue, nil)
	assert.Error(err, "Creating users with same username should return error.")

	// is it escaped?
	uesc, err := user.CreateUser(db, "USERNAME2\"'; --", "hashedPassword", "random", randomValue, nil)
	assert.Nil(err)
	uesc2, err := user.FindUserFromPassword(db, "USERNAME2\"'; --", "hashedPassword")
	assert.Nil(err)
	assert.Equal(uesc.ID, uesc2.ID)

	u2, err := user.FindUserFromPassword(db, "USERNAME", "password")
	helper.PanicOn(err)
	assert.Equal(u2.Username, "USERNAME")
	assert.Equal(u2.ID, u.ID)

	_, err = user.FindUserFromPassword(db, "USERNAME", "password2")
	assert.Error(err)

	u4, err := user.FindUserFromSession(db, user.SessionUser{
		ID:       u.ID,
		Username: "USERNAME",
		Session:  u.Session,
	})
	helper.PanicOn(err)
	assert.Equal(u4.ID, u.ID)
	assert.Equal(u4.Username, u.Username)

	os.Remove("./test.db")
}
