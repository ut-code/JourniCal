package user_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"github.com/ut-code/JourniCal/backend/pkg/random"
	mockctx "github.com/ut-code/JourniCal/backend/pkg/tests/mock-context"
	_ "github.com/ut-code/JourniCal/backend/pkg/tests/run-test-at-root"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type testuser struct {
	username string
	password string
	context  echo.Context
}

func TestRandomTest(t *testing.T) {
	// TODO: this might be dangerous
	options.STATIC_USER = false

	assert := assert.New(t)

	db, err := gorm.Open(sqlite.Open(":memory:"))
	helper.PanicOn(err)
	helper.PanicOn(db.AutoMigrate(&user.User{}))

	count := 5

	var users = make(map[string]testuser)

	for i := range count {
		username := random.String(20)
		password := random.String(32)
		if i == 0 {
			// check empty username
			username = ""
			password = ""
		}

		// check if there is someone with same username
		_, conflict := users[username]

		u, err := user.Create(db, username, password, random.String(20), random.String(20), nil)
		if conflict {
			assert.Error(err)
			continue
		}
		assert.Nil(err)
		c := mockctx.New()
		u.Save(c)

		users[username] = testuser{
			username: username,
			password: password,
			context:  c,
		}
	}

	for _, u := range users {
		u1, err := user.FindUserFromPassword(db, u.username, u.password)
		assert.Nil(err)
		u2, err := user.FromEchoContext(db, u.context)
		assert.Nil(err)
		// cache
		u3, err := user.FromEchoContext(db, u.context)
		assert.Nil(err)
		assert.Equal(u.username, u1.Username)
		assert.Equal(u.username, u2.Username)
		assert.Equal(u.username, u3.Username)
	}
}
