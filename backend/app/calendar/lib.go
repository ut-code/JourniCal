package calendar

import (
	"fmt"
	"os"
	"strings"
	"time"

	"context"

	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func init() {
	if os.Getenv("NO_CREDENTIALS_REQUIRED") == "true" {
		return
	}

}

func CalendarSample(ctx context.Context, config oauth2.Config, tok *oauth2.Token) {
	client := config.Client(ctx, tok)

	// if there is any way to keep the context of a connection between client, (maybe a map[user_id, service]?)
	// service can be cached there.
	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	helper.ErrorLog(err)

	// timezoneTokyo := Timezone{Offset: "+09:00", Area: "Asia/Tokyo"}
	/* CreateEvent(service, "primary", &calendar.Event{
		Summary:  "Hello",
		Location: "World",
		Start:    DateTime("2024-03-16", "13:00:00", timezoneTokyo),
		End:      DateTime("2024-03-16", "17:00:00", timezoneTokyo),
	}) */

	march1 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.Now().Local().Location())
	march31 := time.Date(2024, 3, 31, 23, 59, 59, 999, time.Now().Local().Location())
	evs := GetNEventsForward(service, "primary", march1, 5)
	for _, ev := range evs {
		fmt.Println(prettyFormatEvent(ev))
	}
	fmt.Println("--------------------------")
	evs2 := GetEventsInRange(service, "primary", march1, march31)
	for _, ev := range evs2 {
		fmt.Println(prettyFormatEvent(ev))
	}
}

func prettyFormatEvent(e *calendar.Event) string {
	var attachments_urls []string
	for _, a := range e.Attachments {
		attachments_urls = append(attachments_urls, a.FileUrl+" | "+a.FileId)
	}
	return strings.Join([]string{
		e.Summary,
		e.Description,
		e.Start.Date,
		e.Start.DateTime,
		e.End.Date,
		e.End.DateTime,
		"attachments: " + strings.Join(attachments_urls, " , "),
	}, "|")
}

// this operation halts the app if there is no credentials.json found.
func ReadCredentials() *oauth2.Config {
	bytes, err := os.ReadFile("credentials.json")
	helper.ErrorLog(err, "Failed reading credentials.json")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	helper.ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}

type Timezone struct {
	Offset string
	Area   string
}

type calendar_id = string

// creates event with time specified. for all-day events, use CreateAllDayEvent (might not be implemented yet).
func CreateEvent(service *calendar.Service, calendar_id calendar_id, evt *calendar.Event) {
	_, err := service.Events.Insert(calendar_id, evt).Do()
	helper.ErrorLog(err, "Unable to create event")
}

func RFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func GetNEventsForward(service *calendar.Service, calendar_id calendar_id, start time.Time, count int) []*calendar.Event {
	events, err := service.Events.List(calendar_id).ShowDeleted(false).SingleEvents(false).TimeMin(start.Format(time.RFC3339)).MaxResults(int64(count) + 1).Do()
	helper.ErrorLog(err, "Getting Calendar Events Failed in function GetNEventsForward()")
	return events.Items
}

func GetEventsInRange(service *calendar.Service, calendar_id calendar_id, start time.Time, end time.Time) []*calendar.Event {
	events, err := service.Events.List(calendar_id).SingleEvents(true).TimeMin(start.Format(time.RFC3339)).TimeMax(end.Format(time.RFC3339)).Do()
	helper.ErrorLog(err, "Getting Calendar Events Failed in function GetEventsInRange()")
	return events.Items
}

func DateTime(date, time string, timezone Timezone) *calendar.EventDateTime {
	return &calendar.EventDateTime{
		DateTime: date + "T" + time + timezone.Offset,
		TimeZone: timezone.Area,
	}
}

func AllDayDateTime(date string, timezone Timezone) *calendar.EventDateTime {
	return &calendar.EventDateTime{
		Date:     date,
		TimeZone: timezone.Area,
	}
}
