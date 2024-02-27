package models

import "time"

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
