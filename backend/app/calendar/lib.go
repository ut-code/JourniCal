package calendar

import (
	"time"

	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"google.golang.org/api/calendar/v3"
)

type calendar_id = string

// creates event with time specified. for all-day events, use CreateAllDayEvent (might not be implemented yet).
func CreateEvent(service *calendar.Service, calendar_id calendar_id, evt *calendar.Event) {
	_, err := service.Events.Insert(calendar_id, evt).Do()
	helper.ErrorLog(err, "Unable to create event")
}

func GetEventsInRange(service *calendar.Service, calendar_id calendar_id, start time.Time, end time.Time) []*calendar.Event {
	events, err := service.Events.List(calendar_id).SingleEvents(true).TimeMin(start.Format(time.RFC3339)).TimeMax(end.Format(time.RFC3339)).Do()
	helper.ErrorLog(err, "Getting Calendar Events Failed in function GetEventsInRange()")
	return events.Items
}
