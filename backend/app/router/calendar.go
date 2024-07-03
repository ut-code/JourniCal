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
			return err
		}
		start, err := strconv.Atoi(c.Param("start_unix"))
		if err != nil {
			c.String(400, "Bad request: invalid start time")
			return err
		}
		end, err := strconv.Atoi(c.Param("end_unix"))
		if err != nil {
			c.String(400, "Bad request: invalid end time")
		}
		return c.JSON(http.StatusOK, calendar.GetEventsInRange(srv, "primary", time.Unix(int64(start), 0), time.Unix(int64(end), 0)))
	})
}
