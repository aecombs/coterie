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
func GetChapters(chapterTable *models.ChapterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		chapters, err := chapterTable.ChaptersLister()
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(chapters)
	}
}

//Show
func GetChapter(chapterTable *models.ChapterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		chapterID := chi.URLParam(r, "chapterID")

		chapter, err := chapterTable.ChapterGetter(chapterID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(chapter)
	}
}

//Create
func AddChapter(chapterTable *models.ChapterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		orgID, _ := strconv.Atoi(body["organization_id"])
		chapter := models.Chapter{
			// Text:           body["text"],
			// Date:           body["date"],
			// OrganizationID: orgID,
			// CreatedAt:      time.Now().String(),
			// UpdatedAt:      time.Now().String(),
		}

		result, err := chapterTable.ChapterAdder(chapter)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateChapter(chapterTable *models.ChapterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		chapterID := chi.URLParam(r, "chapterID")
		body := map[string]string{}
		req.BindBody(&body)

		chapID, _ := strconv.Atoi(chapterID)
		chapter := models.Chapter{
			ID: chapID,
			// Text:      body["text"],
			// Date:      body["date"],
			UpdatedAt: time.Now().String(),
		}

		result, err := chapterTable.ChapterUpdater(chapter)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Delete
func DeleteChapter(chapterTable *models.ChapterTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		chapterID := chi.URLParam(r, "chapterID")

		err := chapterTable.ChapterDeleter(chapterID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
