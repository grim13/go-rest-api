package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var events = []Event{}

func (e *Event) Save() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
