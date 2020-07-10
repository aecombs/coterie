package controllers

import (
	"coterie/platform/models"
	"net/http"


	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

type announcementsResource struct{}

// Routes creates a REST router for the announcements resource
func (rs announcementsResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.Index)   // GET /announcements - read a list of announcements
	r.Post("/", rs.Create) // POST /announcements - create a new todo and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a announcements map, and lets actually load/manipulate
		r.Get("/", rs.Show)      // GET /announcements/{id} - read a single todo by :id
		r.Put("/", rs.Update)    // PUT /announcements/{id} - update a single todo by :id
		r.Delete("/", rs.Delete) // DELETE /announcements/{id} - delete a single todo by :id
	})

	return r
}

func (rs announcementsResource) Index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("aaa list of stuff.."))
	res, _ := yin.Event(w, r)
	announcements := models.ListAllAnnouncements()
	res.SendJSON(announcements)
}

func (rs announcementsResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}

func (rs announcementsResource) Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa show"))
}

func (rs announcementsResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa update"))
}

func (rs announcementsResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa delete"))
}
