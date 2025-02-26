package api

import (
	"dev11/internal/models"
	"dev11/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// ResultResponse - result response struct
type ResultResponse struct {
	Result []models.Event `json:"result"`
}

// ErrorResponse - error response struct
type ErrorResponse struct {
	Err string `json:"error"`
}
type Handler struct {
	eventService EventService
}

// NewHandler - creates new handler instance
func NewHandler() *Handler {
	return &Handler{
		eventService: service.New(),
	}
}

type EventService interface {
	CreateEvent(event *models.Event) error
	UpdateEvent(userID, eventID int, newEvent *models.Event) error
	DeleteEvent(userID, eventID int) error
	EventsForDay(userId string, day time.Time) ([]models.Event, error)
	EventsForWeek(userId string, week int) ([]models.Event, error)
	EventsForMonth(userId string, month time.Month) ([]models.Event, error)
}

func (h *Handler) RouteInit(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.CreateEventHandler)
	mux.HandleFunc("/update_event", h.UpdateEventHandler)
	mux.HandleFunc("/delete_event", h.DeleteEventHandler)
	mux.HandleFunc("/events_for_day/", h.EventsForDayHandler)
	mux.HandleFunc("/events_for_week/", h.EventsForWeekHandler)
	mux.HandleFunc("/events_for_month/", h.EventsForMonthHandler)
}

func (h *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	err = h.eventService.CreateEvent(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.resultResponse(w, []models.Event{event})
	log.Printf("event for user %d with id %d created successfully", event.UserID, event.EventID)
}

func (h *Handler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	err = h.eventService.UpdateEvent(event.UserID, event.EventID, &event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.resultResponse(w, []models.Event{event})
	log.Printf("event for user %d with id %d updated successfully", event.UserID, event.EventID)

}
func (h *Handler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.eventService.DeleteEvent(event.UserID, event.EventID)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.resultResponse(w, []models.Event{event})
	log.Printf("event for user %d with id %d deleted successfully", event.UserID, event.EventID)

}
func (h *Handler) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	events, err := h.eventService.EventsForDay(userID, t)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.resultResponse(w, events)
	log.Printf("events for %d.%d.%d were shown successfully", t.Day(), t.Month(), t.Year())

}
func (h *Handler) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	_, week := t.ISOWeek()

	events, err := h.eventService.EventsForWeek(userID, week)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.resultResponse(w, events)
	log.Printf("events for week %d year %d were shown successfully", week, t.Year())

}
func (h *Handler) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusBadRequest, err)
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	events, err := h.eventService.EventsForMonth(userID, t.Month())
	if err != nil {
		log.Printf("status: %d, error: %s", http.StatusServiceUnavailable, err)
		h.errorResponse(w, err, http.StatusServiceUnavailable)
	}
	h.resultResponse(w, events)
	log.Printf("events for month %d year %d were shown successfully", t.Month(), t.Year())

}

// ResultResponse - positive response
func (h *Handler) resultResponse(w http.ResponseWriter, events []models.Event) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.MarshalIndent(&ResultResponse{Result: events}, " ", "")
	_, err := w.Write(result)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// ErrorResponse - response with error status
func (h *Handler) errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
	http.Error(w, string(jsonErr), status)
}

// 2015-09-15T14:00:12Z
// events_for_day/1/2015-09-15T14:00:12Z
