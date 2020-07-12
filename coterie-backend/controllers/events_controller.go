package controllers

import (
	"coterie/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//Index
func GetEvents(eventTable *models.EventTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		events, err := eventTable.EventsLister()
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(events)
	}
}

//Show
func GetEvent(eventTable *models.EventTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		eventID := chi.URLParam(r, "eventID")

		event, err := eventTable.EventGetter(eventID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(event)
	}
}

//Create
func AddEvent(eventTable *models.EventTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(body["organization_id"])
		event := models.Event{
			Name:           body["name"],
			Occasion:       body["occasion"],
			Date:           body["date"],
			Description:    body["description"],
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := eventTable.EventAdder(event)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateEvent(eventTable *models.EventTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		eventID := chi.URLParam(r, "eventID")
		body := map[string]string{}
		req.BindBody(&body)

		evID, _ := strconv.Atoi(eventID)
		event := models.Event{
			ID:          evID,
			Name:        body["name"],
			Occasion:    body["occasion"],
			Date:        body["date"],
			Description: body["description"],
			UpdatedAt:   time.Now().String(),
		}

		result, err := eventTable.EventUpdater(event)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Delete
func DeleteEvent(eventTable *models.EventTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		eventID := chi.URLParam(r, "eventID")

		err := eventTable.EventDeleter(eventID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
