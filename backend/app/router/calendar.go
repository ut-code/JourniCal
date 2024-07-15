package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ut-code/JourniCal/backend/app/calendar"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func Calendar(g *echo.Group, db *gorm.DB) {
	g.GET("/get-events-in-range/:start_unix/:end_unix", func(c echo.Context) error {
		srv, err := calendar.SrvFromContext(db, c)
		if err != nil {
			return c.String(500, "Couldn't start server")
		}
		start, err := strconv.Atoi(c.Param("start_unix"))
		if err != nil {
			return c.String(400, "Bad request: invalid start time")
		}
		end, err := strconv.Atoi(c.Param("end_unix"))
		if err != nil {
			return c.String(400, "Bad request: invalid end time")
		}
		evs, err := calendar.GetEventsInRange(srv, "primary", time.Unix(int64(start), 0), time.Unix(int64(end), 0))
		if err != nil {
			return c.String(500, "Couldn't fetch events")
		}
		return c.JSON(http.StatusOK, evs)
	})
}
