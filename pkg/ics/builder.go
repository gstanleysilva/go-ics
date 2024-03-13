package ics

import (
	"bytes"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMissingParameter = errors.New("missing parameter")
)

type CalendarData struct {
	Title       string
	Start       time.Time
	End         time.Time
	UID         *string
	Description *string
	Location    *string
	Organizer   *string
	Status      *string
	Priority    *string
}

func (d CalendarData) Validate() error {
	if d.Title == "" || d.Start.IsZero() || d.End.IsZero() {
		return ErrMissingParameter
	}
	return nil
}

type Builder struct {
	data []CalendarData
}

func NewBuilder(data []CalendarData) (*Builder, error) {

	for _, data := range data {
		if err := data.Validate(); err != nil {
			return nil, err
		}
	}

	return &Builder{
		data: data,
	}, nil

}

func (b *Builder) getEvents() []Event {
	events := make([]Event, 0, len(b.data))

	for _, data := range b.data {
		event := NewEvent()

		event.AddString("SUMMARY", data.Title)
		event.AddDate("DTSTART", data.Start)
		event.AddDate("DTEND", data.End)
		event.AddDate("DTSTAMP", time.Now())

		if data.UID == nil {
			event.AddString("UID", uuid.NewString())
		} else {
			event.AddString("UID", *data.UID)
		}
		if data.Description != nil {
			event.AddString("DESCRIPTION", *data.Description)
		}
		if data.Location != nil {
			event.AddString("LOCATION", *data.Location)
		}
		if data.Organizer != nil {
			event.AddString("ORGANIZER", *data.Organizer)
		}
		if data.Status != nil {
			event.AddString("STATUS", *data.Status)
		}
		if data.Priority != nil {
			event.AddString("PRIORITY", *data.Priority)
		}

		events = append(events, event)
	}

	return events
}

func (b *Builder) Build() *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})

	body := NewBody(Version2_0, MethodRequest)

	for _, event := range b.getEvents() {
		body.AddEvent(event)
	}

	NewEncoder(buffer).Encode(body)
	return buffer
}
