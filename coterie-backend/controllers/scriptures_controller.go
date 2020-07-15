package controllers

import (
	"coterie/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//GetScriptures gets all the scrips for an org
func GetScriptures(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")

		scriptures, err := scriptureTable.ScripturesLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(scriptures)
	}
}

//Show
func GetScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		scriptureID := chi.URLParam(r, "scriptureID")

		scripture, err := scriptureTable.ScriptureGetter(scriptureID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(scripture)
	}
}

//Create
func AddScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(body["organization_id"])
		scripture := models.Scripture{
			Name:           body["name"],
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := scriptureTable.ScriptureAdder(scripture)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		scriptureID := chi.URLParam(r, "scriptureID")
		body := map[string]string{}
		req.BindBody(&body)

		scripID, _ := strconv.Atoi(scriptureID)
		scripture := models.Scripture{
			ID:        scripID,
			Name:      body["name"],
			UpdatedAt: time.Now().String(),
		}

		result, err := scriptureTable.ScriptureUpdater(scripture)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Delete
func DeleteScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		scriptureID := chi.URLParam(r, "scriptureID")

		err := scriptureTable.ScriptureDeleter(scriptureID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
