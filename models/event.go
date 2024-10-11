package models

import "time"

type Event struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
}

var events = []Event{}

func (e Event) StoreEvents() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events

}
