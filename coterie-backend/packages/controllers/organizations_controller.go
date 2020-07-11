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
func GetOrganizations(organizationTable *models.OrganizationTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		organizations, err := organizationTable.OrganizationsLister()
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(organizations)
	}
}

//Show
func GetOrganization(organizationTable *models.OrganizationTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")

		organization, err := organizationTable.OrganizationGetter(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(organization)
	}
}

//Create
func AddOrganization(organizationTable *models.OrganizationTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		userID, _ := strconv.Atoi(body["user_id"])
		tFunds, _ := strconv.Atoi(body["total_funds"])
		organization := models.Organization{
			Name:             body["name"],
			MissionStatement: body["mission_statement"],
			TotalFunds:       tFunds,
			UserID:           userID,
			CreatedAt:        time.Now().String(),
			UpdatedAt:        time.Now().String(),
		}

		result, err := organizationTable.OrganizationAdder(organization)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateOrganization(organizationTable *models.OrganizationTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(organizationID)
		tFunds, _ := strconv.Atoi(body["total_funds"])
		organization := models.Organization{
			ID:               orgID,
			Name:             body["name"],
			MissionStatement: body["mission_statement"],
			TotalFunds:       tFunds,
			UpdatedAt:        time.Now().String(),
		}

		result, err := organizationTable.OrganizationUpdater(organization)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Delete
func DeleteOrganization(organizationTable *models.OrganizationTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")

		err := organizationTable.OrganizationDeleter(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
