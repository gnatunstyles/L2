package models

import (
	"time"
)

type Event struct {
	EventID int       `json:"event_id"`
	UserID  int       `json:"user_id"`
	Title   string    `json:"title"`
	Info    string    `json:"info"`
	Date    time.Time `json:"date"`
}
