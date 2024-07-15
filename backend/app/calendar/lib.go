package calendar

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"context"

	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

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
	evs, err := GetNEventsForward(service, "primary", march1, 5)
	helper.ErrorLog(err)
	for _, ev := range evs {
		fmt.Println(prettyFormatEvent(ev))
	}
	fmt.Println("--------------------------")
	evs2, err := GetEventsInRange(service, "primary", march1, march31)
	helper.ErrorLog(err)
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

type Timezone struct {
	Offset string
	Area   string
}

type calendar_id = string

// creates event with time specified. for all-day events, use CreateAllDayEvent (might not be implemented yet).
func CreateEvent(service *calendar.Service, calendar_id calendar_id, evt *calendar.Event) error {
	_, err := service.Events.Insert(calendar_id, evt).Do()
	return errors.New("Unable to create event: " + err.Error())
}

func RFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func GetNEventsForward(service *calendar.Service, calendar_id calendar_id, start time.Time, count int) ([]*calendar.Event, error) {
	events, err := service.Events.List(calendar_id).ShowDeleted(false).SingleEvents(false).TimeMin(start.Format(time.RFC3339)).MaxResults(int64(count) + 1).Do()
	return events.Items, err
}

func GetEventsInRange(service *calendar.Service, calendar_id calendar_id, start time.Time, end time.Time) ([]*calendar.Event, error) {
	events, err := service.Events.List(calendar_id).SingleEvents(true).TimeMin(start.Format(time.RFC3339)).TimeMax(end.Format(time.RFC3339)).Do()
	return events.Items, err
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
