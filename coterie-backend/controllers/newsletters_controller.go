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

//GetNewsletters retrieves all newsletters from the DB for a given org
func GetNewsletters(newsletterTable *models.NewsletterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		organizationID := chi.URLParam(r, "organizationID")

		newsletters, err := newsletterTable.NewslettersLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(newsletters)
	}
}

//AddNewsletter adds a new newsletter to the DB for the appropriate org
func AddNewsletter(newsletterTable *models.NewsletterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)

		organizationID := chi.URLParam(r, "organizationID")
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(organizationID)
		newsletter := models.Newsletter{
			Header:         body["header"],
			Description:    body["description"],
			Date:           body["date"],
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := newsletterTable.NewsletterAdder(newsletter)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//UNNESTED ROUTES

//GetNewsletter retrieves a single instance of newsletter from the DB
func GetNewsletter(newsletterTable *models.NewsletterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		newsletterID := chi.URLParam(r, "newsletterID")

		newsletter, err := newsletterTable.NewsletterGetter(newsletterID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(newsletter)
	}
}

//UpdateNewsletter updates an newsletter in the DB
func UpdateNewsletter(newsletterTable *models.NewsletterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)

		newsletterID := chi.URLParam(r, "newsletterID")
		body := map[string]string{}
		req.BindBody(&body)

		annID, _ := strconv.Atoi(newsletterID)
		newsletter := models.Newsletter{
			ID:          annID,
			Header:      body["header"],
			Description: body["description"],
			Date:        body["date"],
			UpdatedAt:   time.Now().String(),
		}

		result, err := newsletterTable.NewsletterUpdater(newsletter)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//DeleteNewsletter removes an newsletter from the DB
func DeleteNewsletter(newsletterTable *models.NewsletterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		newsletterID := chi.URLParam(r, "newsletterID")

		err := newsletterTable.NewsletterDeleter(newsletterID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
