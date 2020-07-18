package controllers

import (
	"coterie/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//NESTED

//GetHolidays will get all the holidays for a given org
func GetHolidays(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		organizationID := chi.URLParam(r, "organizationID")

		holidays, err := holidayTable.HolidaysLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(holidays)
	}
}

//AddHoliday is create action for a specific org
func AddHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		EnableCors(&w)
		body := map[string]string{}
		req.BindBody(&body)
		organizationID := chi.URLParam(r, "organizationID")

		orgID, _ := strconv.Atoi(organizationID)
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

//UNNESTED

//GetHoliday is show action
func GetHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		holidayID := chi.URLParam(r, "holidayID")

		holiday, err := holidayTable.HolidayGetter(holidayID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(holiday)
	}
}

//UpdateHoliday is update action
func UpdateHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		EnableCors(&w)
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

//DeleteHoliday is destroy action
func DeleteHoliday(holidayTable *models.HolidayTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		holidayID := chi.URLParam(r, "holidayID")

		err := holidayTable.HolidayDeleter(holidayID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
