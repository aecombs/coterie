package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

type AnnouncementsResource struct{}

// Routes creates a REST router for the announcements resource
func (rs AnnouncementsResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.Index) // GET /announcements - read a list of announcements
	// r.Get("/new", rs.New) // GET /announcements/new - get the form for a new announcement
	r.Post("/", rs.Create) // POST /announcements - create a new announcement and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Show) // GET /announcements/{id} - read a single announcement by :id
		// r.Get("/edit", rs.Edit)    // GET /announcements/{id}/edit - get form to edit a single announcement by :id
		r.Put("/", rs.Update)    // PUT /announcements/{id} - update a single announcement by :id
		r.Delete("/", rs.Delete) // DELETE /announcements/{id} - delete a single announcement by :id
	})

	return r
}

func (rs AnnouncementsResource) Index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("aaa list of stuff.."))
	res, _ := yin.Event(w, r)
	// announcements := models.ListAllAnnouncements()
	// res.SendJSON(announcements)
	res.SendJSON("There are no announcements here!")
}

// func (rs AnnouncementsResource) New(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa new"))
// }

func (rs AnnouncementsResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}

func (rs AnnouncementsResource) Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa show"))
}

// func (rs AnnouncementsResource) Edit(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("aaa edit"))
// }

func (rs AnnouncementsResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa update"))
}

func (rs AnnouncementsResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa delete"))
}
