package storage

import "time"

type Event struct {
	ID                       string
	Title                    string
	EventDate                time.Time
	ExpirationDate           time.Time
	Description              string
	UserID                   string
	SecondsUntilNotification int
}
