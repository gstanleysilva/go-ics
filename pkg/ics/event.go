package ics

import "time"

type Property struct {
	Name  string
	Value string
}

type Event struct {
	properties []Property
}

func NewEvent() Event {
	return Event{
		properties: make([]Property, 0),
	}
}

func (e *Event) AddString(name, value string) {
	e.properties = append(e.properties, Property{
		Name:  name,
		Value: value},
	)
}

func (e *Event) AddDate(name string, value time.Time) {
	e.properties = append(e.properties, Property{
		Name:  name,
		Value: value.UTC().Format("20060102T150405Z"),
	})
}
