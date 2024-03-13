package ics

import (
	"fmt"
	"io"
)

const (
	Version2_0    = "2.0"
	MethodRequest = "REQUEST"
	MethodPublish = "PUBLISH"
	MethodReply   = "REPLY"
	MethodAdd     = "ADD"
	MethodCancel  = "CANCEL"
	MethodRefresh = "REFRESH"
	MethodCounter = "COUNTER"
)

type EncoderBody struct {
	Version string
	Method  string
	Event   []Event
}

func NewBody(version, method string) EncoderBody {
	return EncoderBody{
		Version: version,
		Method:  method,
		Event:   make([]Event, 0),
	}
}

func (b *EncoderBody) AddEvent(event Event) {
	b.Event = append(b.Event, event)
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(body EncoderBody) error {

	if len(body.Event) == 0 {
		return fmt.Errorf("no events found")
	}

	//Add Header Properties
	fmt.Fprintf(e.w, "BEGIN:VCALENDAR\r\n")
	fmt.Fprintf(e.w, "%s:%s\r\n", "VERSION", body.Version)

	for _, event := range body.Event {
		//Start Event
		fmt.Fprintf(e.w, "BEGIN:VEVENT\r\n")
		for _, p := range event.properties {
			fmt.Fprintf(e.w, "%s:%s\r\n", p.Name, p.Value)
		}
		fmt.Fprint(e.w, "END:VEVENT\r\n")
		//End Event
	}

	fmt.Fprint(e.w, "END:VCALENDAR\r\n")
	return nil
}
