package router

import (
	"net/http"
	"strconv"

	"github.com/ut-code/JourniCal/backend/app/diary"
	"github.com/ut-code/JourniCal/backend/app/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Diary(g *echo.Group, db *gorm.DB) {
	g.GET("/", func(c echo.Context) error {
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusBadRequest, "authentication error")
		}
		json, err := diary.GetAll(db, u)
		if err != nil {
			return c.String(500, err.Error())
		}
		return c.JSON(200, json)
	})

	g.GET("/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 0 {
			return c.String(http.StatusBadRequest, "invalid id format")
		}
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "authentication error")
		}
		diary, err := diary.Get(db, uint(id), u)
		if err != nil {
			// TODO: fix db error included in here
			return c.String(http.StatusForbidden, err.Error())
		}
		return c.JSON(200, diary)
	})

	// CREATE
	g.POST("/", func(c echo.Context) error {
		d := new(diary.Diary)
		// UNTESTED: I'm not sure how this works or how this should be done. Test this function.
		if err := c.Bind(d); err != nil {
			return c.String(http.StatusBadRequest, "failed to bind diary data")
		}
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "invalid user. consider creating new one")
		}
		diary, err := diary.Create(db, d, u)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, diary)
	})

	// UPDATE
	g.PUT("/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 0 {
			return c.String(http.StatusBadRequest, "Invalid formating of id or negative value")
		}
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid user. create new one")
		}

		var newDiary diary.Diary
		if err := c.Bind(&newDiary); err != nil {
			return c.String(http.StatusBadRequest, "Failed to bind diary")
		}
		diary, err := diary.Update(db, uint(id), &newDiary, u)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update")
		}
		return c.JSON(http.StatusAccepted, diary)
	})

	// DELETE
	g.DELETE("/:id", func(c echo.Context) error {
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Authentication error: user not found")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 0 {
			return c.String(http.StatusBadRequest, "Invalid diary ID")
		}
		err = diary.Delete(db, uint(id), u)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete")
		}
		return c.NoContent(http.StatusAccepted)
	})
}
