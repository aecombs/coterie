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

//GetAnnouncements retrieves all announcements from the DB for a given org
func GetAnnouncements(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")

		announcements, err := announcementTable.AnnouncementsLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(announcements)
	}
}

//AddAnnouncement adds a new announcement to the DB for the appropriate org
func AddAnnouncement(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(organizationID)
		announcement := models.Announcement{
			Text:           body["text"],
			Date:           body["date"],
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := announcementTable.AnnouncementAdder(announcement)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//UNNESTED ROUTES

//GetAnnouncement retrieves a single instance of announcement from the DB
func GetAnnouncement(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		announcementID := chi.URLParam(r, "announcementID")

		announcement, err := announcementTable.AnnouncementGetter(announcementID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(announcement)
	}
}

//UpdateAnnouncement updates an announcement in the DB
func UpdateAnnouncement(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		announcementID := chi.URLParam(r, "announcementID")
		body := map[string]string{}
		req.BindBody(&body)

		annID, _ := strconv.Atoi(announcementID)
		announcement := models.Announcement{
			ID:        annID,
			Text:      body["text"],
			Date:      body["date"],
			UpdatedAt: time.Now().String(),
		}

		result, err := announcementTable.AnnouncementUpdater(announcement)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//DeleteAnnouncement removes an announcement from the DB
func DeleteAnnouncement(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		announcementID := chi.URLParam(r, "announcementID")

		err := announcementTable.AnnouncementDeleter(announcementID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
