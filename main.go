package main

import (
	"fmt"
	"time"

	"github.com/gstanleysilva/go-ics/pkg/ics"
)

func main() {
	location := "MyLocation"
	description := "MyDescription"
	organizer := "MyOrganizer"
	priority := "3"

	data := ics.CalendarData{
		Title:       "MyTitle",
		Start:       time.Now(),
		End:         time.Now().Add(time.Hour * 24),
		Description: &description,
		Location:    &location,
		Organizer:   &organizer,
		Status:      nil,
		Priority:    &priority,
	}

	ical, err := ics.NewBuilder([]ics.CalendarData{data})
	if err != nil {
		panic(err)
	}

	fmt.Println(ical.Build().String())
}
