package main

import (
	"coterie/packages/controllers"
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	flag.Parse()
	//open the database!
	db, err := sql.Open("sqlite3", "./database/coterie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Announcements := models.NewAnnouncementTable(db)
	// chapters := models.NewChapterTable(db)
	// events := models.NewEventTable(db)
	// holidays := models.NewHolidayTable(db)
	// members := models.NewMemberTable(db)
	// organizations := models.NewOrganizationTable(db)
	// scriptures := models.NewScriptureTable(db)
	// users := models.NewUserTable(db)

	r := chi.NewRouter()

	// r.Use(yin.SimpleLogger)

	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	// r.Route("/announcements", func(r chi.Router) {
	// 	r.With(paginate).Get("/", controllers.GetAnnouncements(db))
	// 	r.Post("/", controllers.CreateAnnouncement(db))       // POST /Announcements
	// 	// r.Get("/search", controllers.SearchAnnouncements()) // GET /Announcements/search

	// 	r.Route("/{announcementID}", func(r chi.Router) {
	// 		// r.Use(controllers.AnnouncementCtx)       // Load the *Announcement on the request context
	// 		r.Get("/", controllers.GetAnnouncement(db))       // GET /Announcements/123
	// 		r.Put("/", controllers.UpdateAnnouncement(db))    // PUT /Announcements/123
	// 		r.Delete("/", controllers.DeleteAnnouncement(db)) // DELETE /Announcements/123
	// 	})
	// })

	r.Mount("/announcements", controllers.AnnouncementsResource{}.Routes())
	// r.Mount("/todos", todosResource{}.Routes())

	http.ListenAndServe(":3000", r)
}
