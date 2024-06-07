package router

import (
	"JourniCalBackend/calendar"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Calendar(g *echo.Group) {
	// spec:
	// specify time as unix time.
	g.GET("/get-20-events-forward/:start_unix", func(c echo.Context) error {
		srv, err := calendar.SrvFromContext(c)
		if err != nil {
			return err
		}
		t, err := strconv.Atoi(c.Param("start_unix"))
		if err != nil {
			c.String(http.StatusBadRequest, "get-20-events-forward with invalid second path: not unix time (= number)")
			return err
		}
		evs := calendar.GetNEventsForward(srv, "primary", time.Unix(int64(t), 0), 20)
		return c.JSON(200, evs)
	})
}
