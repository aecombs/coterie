package main

import (
	"coterie/controllers"
	"coterie/models"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

func main() {
	flag.Parse()
	//open the database!
	db, err := sql.Open("sqlite3", "./database/coterie.db")
	if err != nil {
		log.Printf("Unable to access database: %s", err.Error())
		log.Fatal(err)
	}
	defer db.Close()

	announcements := models.NewAnnouncementTable(db)
	chapters := models.NewChapterTable(db)
	events := models.NewEventTable(db)
	holidays := models.NewHolidayTable(db)
	members := models.NewMemberTable(db)
	organizations := models.NewOrganizationTable(db)
	scriptures := models.NewScriptureTable(db)
	users := models.NewUserTable(db)

	r := chi.NewRouter()

	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(yin.SimpleLogger)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//Users
	r.Route("/", func(r chi.Router) {
		//auth handling
		// r.Options("/*", controllers.ApproveCors())
		r.Get("/auth/google", controllers.GoogleLogin())
		r.Get("/auth/google/callback", controllers.GoogleCallback(users))
		r.Get("/logout", controllers.LogoutUser())

		r.Route("/users/{userID}", func(r chi.Router) {
			r.Get("/", controllers.GetUser(users))
			r.Put("/", controllers.UpdateUser(users))

			//Organizations
			r.Route("/organizations", func(r chi.Router) {
				r.Get("/", controllers.GetOrganizations(organizations, users))
				r.Post("/", controllers.AddOrganization(organizations, users))

				r.Route("/{organizationID}", func(r chi.Router) {
					r.Get("/", controllers.GetOrganization(organizations))
					r.Put("/", controllers.UpdateOrganization(organizations))
					r.Delete("/", controllers.DeleteOrganization(organizations))

					//nested announcements
					r.Route("/announcements", func(r chi.Router) {
						r.Get("/", controllers.GetAnnouncements(announcements))
						r.Post("/", controllers.AddAnnouncement(announcements))
					})

					//nested events
					r.Route("/events", func(r chi.Router) {
						r.Get("/", controllers.GetEvents(events))
						r.Post("/", controllers.AddEvent(events))
					})

					//nested holidays
					r.Route("/holidays", func(r chi.Router) {
						r.Get("/", controllers.GetHolidays(holidays))
						r.Post("/", controllers.AddHoliday(holidays))
					})

					//nested scriptures
					r.Route("/scriptures", func(r chi.Router) {
						r.Get("/", controllers.GetScriptures(scriptures))
						r.Post("/", controllers.AddScripture(scriptures))
					})
					//nested members
					r.Route("/members", func(r chi.Router) {
						r.Get("/", controllers.GetMembers(members))
						r.Post("/", controllers.AddMember(members))
					})
				})
			})
		})
	})

	//Announcements
	r.Route("/announcements/{announcementID}", func(r chi.Router) {
		r.Get("/", controllers.GetAnnouncement(announcements))
		r.Put("/", controllers.UpdateAnnouncement(announcements))
		r.Delete("/", controllers.DeleteAnnouncement(announcements))
	})

	//Events
	r.Route("/events/{eventID}", func(r chi.Router) {
		r.Get("/", controllers.GetEvent(events))
		r.Put("/", controllers.UpdateEvent(events))
		r.Delete("/", controllers.DeleteEvent(events))
	})

	//Holidays
	r.Route("/holidays/{holidayID}", func(r chi.Router) {
		r.Get("/", controllers.GetHoliday(holidays))
		r.Put("/", controllers.UpdateHoliday(holidays))
		r.Delete("/", controllers.DeleteHoliday(holidays))
	})

	//Members
	r.Route("/members/{memberID}", func(r chi.Router) {
		r.Get("/", controllers.GetMember(members))
		r.Put("/", controllers.UpdateMember(members))
		r.Delete("/", controllers.DeleteMember(members))
	})

	//Scriptures
	r.Route("/scriptures/{scriptureID}", func(r chi.Router) {
		r.Get("/", controllers.GetScripture(scriptures))
		r.Put("/", controllers.UpdateScripture(scriptures))
		r.Delete("/", controllers.DeleteScripture(scriptures))

		//Nested Chapters
		r.Route("/chapters", func(r chi.Router) {
			r.Get("/", controllers.GetChapters(chapters))
			r.Post("/", controllers.AddChapter(chapters))
		})
	})

	//Chapters
	r.Route("/chapters/{chapterID}", func(r chi.Router) {
		r.Get("/", controllers.GetChapter(chapters))
		r.Put("/", controllers.UpdateChapter(chapters))
		r.Delete("/", controllers.DeleteChapter(chapters))
	})

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
