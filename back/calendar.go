package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	cachedToken := readTokenCache()
	// Authenticate will do:
	// if cache == nil,
	//    call askForToken to get token.
	//    call saveToken to save token for future use.
	// authenticate.
	client := Authenticate(ctx, cfg, cachedToken, askForToken, saveToken)

	// if there is any way to keep the context of a connection between client, (or maybe a map[user_id, service]?)
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
	now := time.Now().Format(time.RFC3339)
	evs := GetNEventsForward(service, "primary", now, 20)
	for _, ev := range evs {
		fmt.Println(prettyFormatEvent(ev))
	}
}
func prettyFormatEvent(e *calendar.Event) string {
	return strings.Join([]string{
		e.Summary,
		e.Description,
		e.Start.Date,
		e.Start.DateTime,
		e.End.Date,
		e.End.DateTime,
	}, "/")
}

// this operation halts the app if there is no credentials.json found.
func ReadCredentials() *oauth2.Config {
	bytes, err := os.ReadFile("credentials.json")
	ErrorLog(err, "Failed reading credentials.json")

	cfg, err := google.ConfigFromJSON(bytes, calendar.CalendarScope)
	ErrorLog(err, "Unable to parse client secret file to config")

	return cfg
}

func askForToken(url authUrl) authCode {
	fmt.Println("Go go the link and get token (token will appear in the query) and paste to terminal", url)

	var code string
	_, err := fmt.Scan(&code)
	ErrorLog(err, "unable to read authorization code")
	return code
}

func saveToken(token *oauth2.Token) {
	file := "./token.json"
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	ErrorLog(err, "Unable to cache oauth token")

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func readTokenCache() *oauth2.Token {
	f, err := os.Open("./token.json")
	if err != nil {
		return nil
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	if err != nil {
		return nil
	}
	return t
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

type RFC3339 = string

func GetNEventsForward(service *calendar.Service, calendar_id calendar_id, start RFC3339, count int) []*calendar.Event {
	events, err := service.Events.List(calendar_id).ShowDeleted(false).SingleEvents(true).TimeMin(start).MaxResults(int64(count)).Do()
	ErrorLog(err, "Getting Calendar Events Failed in function GetNEvents()")
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
