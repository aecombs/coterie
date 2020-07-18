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

//GetScriptures gets all the scrips for an org
func GetScriptures(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		organizationID := chi.URLParam(r, "organizationID")

		scriptures, err := scriptureTable.ScripturesLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(scriptures)
	}
}

//AddScripture is creat action to a single org
func AddScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		EnableCors(&w)
		body := map[string]string{}
		req.BindBody(&body)
		organizationID := chi.URLParam(r, "organizationID")

		orgID, _ := strconv.Atoi(organizationID)
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

//UNNESTED

//GetScripture is show action
func GetScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		scriptureID := chi.URLParam(r, "scriptureID")

		scripture, err := scriptureTable.ScriptureGetter(scriptureID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(scripture)
	}
}

//UpdateScripture is update action
func UpdateScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		EnableCors(&w)
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

//DeleteScripture is destroy action
func DeleteScripture(scriptureTable *models.ScriptureTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		EnableCors(&w)
		scriptureID := chi.URLParam(r, "scriptureID")

		err := scriptureTable.ScriptureDeleter(scriptureID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
