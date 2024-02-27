package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required"json:"name"`
	Description string    `binding:"required"json:"description"`
	Location    string    `binding:"required"json:"location"`
	DateTime    time.Time `binding:"required"json:"dateTime"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var events = []Event{}

func (e Event) Save() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
