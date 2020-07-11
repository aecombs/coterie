package controllers

import (
	"coterie/packages/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//Index
func GetHolidays(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		holidays, err := holidayTable.HolidaysLister()
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(holidays)
	}
}

//Show
func GetHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		holidayID := chi.URLParam(r, "holidayID")

		holiday, err := holidayTable.HolidayGetter(holidayID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(holiday)
	}
}

//Create
func AddHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(body["organization_id"])
		holiday := models.Holiday{
			Name:           body["name"],
			Date:           body["date"],
			Description:    body["description"],
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := holidayTable.HolidayAdder(holiday)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		holidayID := chi.URLParam(r, "holidayID")
		body := map[string]string{}
		req.BindBody(&body)

		holID, _ := strconv.Atoi(holidayID)
		holiday := models.Holiday{
			ID:          holID,
			Name:        body["name"],
			Date:        body["date"],
			Description: body["description"],
			UpdatedAt:   time.Now().String(),
		}

		result, err := holidayTable.HolidayUpdater(holiday)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Delete
func DeleteHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		holidayID := chi.URLParam(r, "holidayID")

		err := holidayTable.HolidayDeleter(holidayID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
