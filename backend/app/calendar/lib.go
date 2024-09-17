package calendar

import (
	"errors"
	"time"

	"google.golang.org/api/calendar/v3"
)

type calendar_id = string

// creates event with time specified. for all-day events, use CreateAllDayEvent (might not be implemented yet).
func CreateEvent(service *calendar.Service, calendar_id calendar_id, evt *calendar.Event) error {
	_, err := service.Events.Insert(calendar_id, evt).Do()
	return errors.New("Unable to create event: " + err.Error())
}

func GetEventsInRange(service *calendar.Service, calendar_id calendar_id, start time.Time, end time.Time) ([]*calendar.Event, error) {
	events, err := service.Events.List(calendar_id).SingleEvents(true).TimeMin(start.Format(time.RFC3339)).TimeMax(end.Format(time.RFC3339)).Do()
	return events.Items, err
}
