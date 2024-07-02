package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/app"
	"github.com/ut-code/JourniCal/backend/pkg/curl"
	"google.golang.org/api/calendar/v3"
)

func init() {
	go app.Serve(3000)
	time.Sleep(1) // wait for the server to start
}

/* -------------------- */

func TestFull(t *testing.T) {

	// TODO: read README and fix this ci test
	if true {
		return
	}
	assert := assert.New(t)

	// assert := assertion.New(t)
	curl := curl.WithCookie("./token.cookie")
	curl.PrefixPath("localhost:3000")
	local := time.Now().Local().Location()
	var events []calendar.Event
	err := curl.JSON(GetEventsInRange(time.Date(2024, 4, 1, 0, 0, 0, 0, local), time.Date(2024, 5, 1, 0, 0, 0, 0, local)), &events)
	assert.Nil(err, "err on curl")
	assert.Equal(len(events), 1)
	if len(events) == 0 {
		panic("len(events) == 0")
	}
	assert.True(events[0].Start.Date == "2024-04-01")
	assert.True(events[0].Summary == "First Event of April. Used in Go Test, DO NOT CHANGE")
}

func GetEventsInRange(start time.Time, end time.Time) string {
	return "/api/calendar/get-events-in-range/" + fmt.Sprint(start.Unix()) + "/" + fmt.Sprint(end.Unix())
}
