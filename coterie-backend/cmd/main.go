package main

import (
	"coterie/packages/controllers"
	"coterie/packages/models"
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

func main() {
	flag.Parse()
	//open the database!
	db, err := sql.Open("sqlite3", "./database/coterie.db")
	if err != nil {
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
	// r.Use(middleware.Logger)
	r.Use(yin.SimpleLogger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	//Announcements
	r.Get("/announcements", controllers.GetAnnouncements(announcements))
	r.Get("/announcements/{announcementID}", controllers.GetAnnouncement(announcements))
	r.Post("/announcements", controllers.AddAnnouncement(announcements))
	r.Put("/announcements/{announcementID}", controllers.UpdateAnnouncement(announcements))
	r.Delete("/announcements/{announcementID}", controllers.DeleteAnnouncement(announcements))

	//Chapters
	r.Get("/chapters", controllers.GetChapters(chapters))
	r.Get("/chapters/{chapterID}", controllers.GetChapter(chapters))
	r.Post("/chapters", controllers.AddChapter(chapters))
	r.Put("/chapters/{chapterID}", controllers.UpdateChapter(chapters))
	r.Delete("/chapters/{chapterID}", controllers.DeleteChapter(chapters))

	//Events
	r.Get("/events", controllers.GetEvents(events))
	r.Get("/events/{eventID}", controllers.GetEvent(events))
	r.Post("/events", controllers.AddEvent(events))
	r.Put("/events/{eventID}", controllers.UpdateEvent(events))
	r.Delete("/events/{eventID}", controllers.DeleteEvent(events))

	//Holidays
	r.Get("/holidays", controllers.GetHolidays(holidays))
	r.Get("/holidays/{holidayID}", controllers.GetHoliday(holidays))
	r.Post("/holidays", controllers.AddHoliday(holidays))
	r.Put("/holidays/{holidayID}", controllers.UpdateHoliday(holidays))
	r.Delete("/holidays/{holidayID}", controllers.DeleteHoliday(holidays))

	//Members
	r.Get("/members", controllers.GetMembers(members))
	r.Get("/members/{memberID}", controllers.GetMember(members))
	r.Post("/members", controllers.AddMember(members))
	r.Put("/members/{memberID}", controllers.UpdateMember(members))
	r.Delete("/members/{memberID}", controllers.DeleteMember(members))

	//Organizations
	r.Get("/organizations", controllers.GetOrganizations(organizations))
	r.Get("/organizations/{organizationID}", controllers.GetOrganization(organizations))
	r.Post("/organizations", controllers.AddOrganization(organizations))
	r.Put("/organizations/{organizationID}", controllers.UpdateOrganization(organizations))
	r.Delete("/organizations/{organizationID}", controllers.DeleteOrganization(organizations))

	//Scriptures
	r.Get("/scriptures", controllers.GetScriptures(scriptures))
	r.Get("/scriptures/{scriptureID}", controllers.GetScripture(scriptures))
	r.Post("/scriptures", controllers.AddScripture(scriptures))
	r.Put("/scriptures/{scriptureID}", controllers.UpdateScripture(scriptures))
	r.Delete("/scriptures/{scriptureID}", controllers.DeleteScripture(scriptures))

	http.ListenAndServe(":3000", r)
}
