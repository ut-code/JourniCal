package router

import (
	"net/http"
	"strconv"

	"github.com/ut-code/JourniCal/backend/app/journal"
	"github.com/ut-code/JourniCal/backend/app/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Journal(g *echo.Group, db *gorm.DB) {
	g.GET("/", func(c echo.Context) error {
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusBadRequest, "authentication error")
		}
		json, err := journal.GetAll(db, u)
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
		journal, err := journal.Get(db, uint(id), u)
		if err != nil {
			// TODO: fix db error included in here
			return c.String(http.StatusForbidden, err.Error())
		}
		return c.JSON(200, journal)
	})

	// CREATE
	g.POST("/", func(c echo.Context) error {
		d := new(journal.Journal)
		// UNTESTED: I'm not sure how this works or how this should be done. Test this function.
		if err := c.Bind(d); err != nil {
			return c.String(http.StatusBadRequest, "failed to bind journal data")
		}
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "invalid user. consider creating new one")
		}
		err = journal.Create(db, d, u)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, d)
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

		var newJournal journal.Journal
		if err := c.Bind(&newJournal); err != nil {
			return c.String(http.StatusBadRequest, "Failed to bind journal")
		}
		journal, err := journal.Update(db, uint(id), &newJournal, u)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update")
		}
		return c.JSON(http.StatusAccepted, journal)
	})

	// DELETE
	g.DELETE("/:id", func(c echo.Context) error {
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Authentication error: user not found")
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 0 {
			return c.String(http.StatusBadRequest, "Invalid journal ID")
		}
		err = journal.Delete(db, uint(id), u)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete")
		}
		return c.NoContent(http.StatusAccepted)
	})
}
