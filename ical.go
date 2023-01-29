package main

import (
	ics "github.com/arran4/golang-ical"
	"net/http"
	"strings"
	"time"
)

type calendarEvent struct {
	Summary     string
	Location    string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

func ParseEvent(event *ics.VEvent) calendarEvent {
	startTime, err := event.GetStartAt()
	if err != nil {
		panic(err)
	}
	endTime, err := event.GetEndAt()
	if err != nil {
		panic(err)
	}
	return calendarEvent{
		Summary:     event.GetProperty("SUMMARY").Value,
		Location:    event.GetProperty("LOCATION").Value,
		Description: strings.TrimSpace(event.GetProperty("DESCRIPTION").Value),
		StartTime:   startTime,
		EndTime:     endTime,
	}
}

func getEvents(calendarId string) []calendarEvent {
	resp, err := http.Get("https://edt.telecomnancy.univ-lorraine.fr/" + calendarId + ".ics")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	calendar, err := ics.ParseCalendar(resp.Body)
	if err != nil {
		panic(err)
	}
	var events []calendarEvent
	for _, event := range calendar.Events() {
		events = append(events, ParseEvent(event))
	}
	return events
}
