package main

import (
	"fmt"
	"os"
	"strings"

	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func CalendarSample() {
	// do this at top-level initialization in main (or, copy & paste from inside the function)
	ctx := context.Background()
	cfg := ReadCredentials()

	code, err := ReadFile("./saved-code.txt") // read token from cookie in the product code.
	if err != nil || code == "" {
		authURL := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Println("Go go the link and get token (token will appear in the query) and paste to terminal", authURL)

		var code string
		_, err := fmt.Scan(&code)
		ErrorLog(err, "unable to read authorization code")

		f, err := os.Create("./saved-code.txt")
		ErrorLog(err)
		_, err = f.Write([]byte(code))
		ErrorLog(err)

		tok, err := cfg.Exchange(ctx, code)
		ErrorLog(err, "Unable to retrieve token from web")

		SaveToken(code, tok)
	}

	tok, err := ReadToken(ctx, code, cfg)
	ErrorLog(err)
	client := cfg.Client(ctx, tok)

	// if there is any way to keep the context of a connection between client, (maybe a map[user_id, service]?)
	// service can be cached there.
	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	ErrorLog(err, "Failed to create service")

	// timezoneTokyo := Timezone{Offset: "+09:00", Area: "Asia/Tokyo"}
	/* CreateEvent(service, "primary", &calendar.Event{
		Summary:  "Hello",
		Location: "World",
		Start:    DateTime("2024-03-16", "13:00:00", timezoneTokyo),
		End:      DateTime("2024-03-16", "17:00:00", timezoneTokyo),
	}) */

	evs := GetNEventsForward(service, "primary", RFC3339("2024-03-01T00:00:00+09:00"), 5)
	for _, ev := range evs {
		fmt.Println(prettyFormatEvent(ev))
	}
	fmt.Println("--------------------------")
	evs2 := GetEventsInRange(service, "primary", RFC3339("2024-03-01T00:00:00+09:00"), RFC3339("2024-03-31T23:59:59+09:00"))
	for _, ev := range evs2 {
		fmt.Println(prettyFormatEvent(ev))
	}
}
func prettyFormatEvent(e *calendar.Event) string {
	var attachments_urls []string
	for _, a := range e.Attachments {
		attachments_urls = append(attachments_urls, a.FileUrl)
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
	ErrorLog(err, "Failed reading credentials.json")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	ErrorLog(err, "Unable to parse client secret file to config")

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
	ErrorLog(err, "Unable to create event")
}

type RFC3339 string

func GetNEventsForward(service *calendar.Service, calendar_id calendar_id, start RFC3339, count int) []*calendar.Event {
	events, err := service.Events.List(calendar_id).ShowDeleted(false).SingleEvents(false).TimeMin(string(start)).MaxResults(int64(count) + 1).Do()
	ErrorLog(err, "Getting Calendar Events Failed in function GetNEventsForward()")
	return events.Items
}
func GetEventsInRange(service *calendar.Service, calendar_id calendar_id, start RFC3339, end RFC3339) []*calendar.Event {
	events, err := service.Events.List(calendar_id).SingleEvents(false).TimeMin(string(start)).TimeMax(string(end)).Do()
	ErrorLog(err, "Getting Calendar Events Failed in function GetEventsInRange()")
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
