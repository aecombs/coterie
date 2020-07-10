package controllers

import (
	"coterie/packages/models"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

// type AnnouncementsResource struct{}

// Routes creates a REST router for the announcements resource
// func (rs AnnouncementsResource) Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Use(yin.SimpleLogger)

// 	r.Get("/", rs.Index) // GET /announcements - read a list of announcements
// 	// r.Get("/new", rs.New) // GET /announcements/new - get the form for a new announcement
// 	r.Post("/", rs.Create) // POST /announcements - create a new announcement and persist it

// 	r.Route("/{id}", func(r chi.Router) {
// 		r.Get("/", rs.Show) // GET /announcements/{id} - read a single announcement by :id
// 		// r.Get("/edit", rs.Edit)    // GET /announcements/{id}/edit - get form to edit a single announcement by :id
// 		r.Put("/", rs.Update)    // PUT /announcements/{id} - update a single announcement by :id
// 		r.Delete("/", rs.Delete) // DELETE /announcements/{id} - delete a single announcement by :id
// 	})

// 	return r
// }

//Index
func GetAnnouncements(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)

		announcements, err := announcementTable.AnnouncementsLister()
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(announcements)
	}
}

//Show
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

//Create
func AddAnnouncement(announcementTable *models.AnnouncementTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		announcement := models.Announcement{
			Text:      body["text"],
			Date:      body["date"],
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		result := announcementTable.AnnouncementAdder(announcement)

		result, err := announcementTable.AnnouncementAdder(announcement)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
	}
}

// func (rs AnnouncementsResource) Index(w http.ResponseWriter, r *http.Request) {
// 	res, _ := yin.Event(w, r)
// 	// announcements := models.ListAllAnnouncements()
// 	// res.SendJSON(announcements)
// 	res.SendJSON("There are no announcements here!")
// }

// func (rs AnnouncementsResource) New(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa new"))
// }

// func (rs AnnouncementsResource) Create(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("aaa create"))
// 	res, req := yin.Event(w, r)
// 	body := map[string]string{}
// 	req.BindBody(&body)

// fmt.Println(body)

// res.SendStatus(http.StatusNoContent)

// 	announcement := models.Announcement{
// 		Text: body["text"],
// 		Date: body["date"],
// 		created_at: time.Now(),
// 		updated_at: time.Now(),
// 	}
// 	// feed.Add(item)
// 	res.SendJSON(announcement)

// }

// func (rs AnnouncementsResource) Show(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa show"))
// }

// func (rs AnnouncementsResource) Edit(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa edit"))
// }

// func (rs AnnouncementsResource) Update(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa update"))
// }

// func (rs AnnouncementsResource) Delete(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa delete"))
// }
