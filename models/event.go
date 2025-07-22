package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// later: adde it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
