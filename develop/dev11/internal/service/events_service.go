package service

import (
	"dev11/internal/models"
	"fmt"
	"strings"
	"time"
)

type Events interface {
	CreateEvent(event *models.Event) error
	UpdateEvent(userID, eventID int, newEvent *models.Event) error
	DeleteEvent(userID, eventID int)
}

type EventsDB struct {
	db map[string]models.Event
}

func New() *EventsDB {
	return &EventsDB{
		db: make(map[string]models.Event),
	}
}

func (e *EventsDB) CreateEvent(event *models.Event) error {
	saveId := fmt.Sprintf("%d:%d", event.UserID, event.EventID)
	if _, ok := e.db[saveId]; !ok {
		e.db[saveId] = *event
		return nil
	} else {
		return fmt.Errorf("error: event with this ID already exists")
	}
}
func (e *EventsDB) UpdateEvent(userID, eventID int, newEvent *models.Event) error {
	saveId := fmt.Sprintf("%d:%d", userID, eventID)
	if _, ok := e.db[saveId]; ok {
		e.db[saveId] = *newEvent
		return nil
	} else {
		return fmt.Errorf("error: event with this ID does not exist")
	}
}
func (e *EventsDB) DeleteEvent(userID, eventID int) error {
	saveId := fmt.Sprintf("%d:%d", userID, eventID)
	if _, ok := e.db[saveId]; ok {
		delete(e.db, saveId)
		return nil
	} else {
		return fmt.Errorf("error: event with this ID does not exist")
	}
}

func (e *EventsDB) EventsForDay(userId string, day time.Time) ([]models.Event, error) {
	var events []models.Event
	for k, v := range e.db {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			if v.Date.Day() == day.Day() {
				events = append(events, v)
			}
		}
	}
	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
func (e *EventsDB) EventsForWeek(userId string, week int) ([]models.Event, error) {
	var events []models.Event
	for k, v := range e.db {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			_, w := v.Date.ISOWeek()
			if w == week {
				events = append(events, v)
			}
		}
	}
	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
func (e *EventsDB) EventsForMonth(userId string, month time.Month) ([]models.Event, error) {
	var events []models.Event
	for k, v := range e.db {
		keyId := strings.Split(k, ":")[0]
		if userId == keyId {
			if v.Date.Month() == month {
				events = append(events, v)
			}
		}
	}
	if len(events) > 0 {
		return events, nil
	} else {
		return nil, fmt.Errorf("error: no events for this date ")
	}
}
