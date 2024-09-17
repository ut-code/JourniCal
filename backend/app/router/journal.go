package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/journal"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/list"
	gcal "google.golang.org/api/calendar/v3"

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

	// start and end must be encoded as string of Unix timestamp
	g.GET("/in-range/:start/:end", func(c echo.Context) error {
		csrv, err := calendar.SrvFromContext(db, c)
		if err != nil {
			return c.String(500, "Couldn't start service")
		}
		user, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(401, "couldn't find user")
		}

		start, err := strconv.Atoi(c.Param("start"))
		if err != nil {
			return c.String(400, "Bad request: invalid start time")
		}
		end, err := strconv.Atoi(c.Param("end"))
		if err != nil {
			return c.String(400, "Bad request: invalid end time")
		}

		events, err := calendar.GetEventsInRange(csrv, "primary", time.Unix(int64(start), 0), time.Unix(int64(end), 0))
		if err != nil {
			return c.String(500, "Couldn't fetch events")
		}

		it := list.ConcurrentMap(events, func(event *gcal.Event) *journal.Journal {
			jnl, err := journal.GetByEvent(db, event.Id, user)
			if err == nil {
				return jnl
			}
			return nil
		})

		result := list.FilterNil(it)
		return c.JSON(200, result)
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

	// GET BY EVENT
	g.GET("/event/:eventID", func(c echo.Context) error {
		eventID := c.Param("eventID")
		u, err := user.FromEchoContext(db, c)
		if err != nil {
			return c.String(http.StatusUnauthorized, "authentication error")
		}
		journal, err := journal.GetByEvent(db, eventID, u)
		if err != nil {
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
