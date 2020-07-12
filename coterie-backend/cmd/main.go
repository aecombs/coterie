package main

import (
	"coterie/controllers"
	"coterie/models"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var mySigningKey = goDotEnvVariable("MY_JWT_TOKEN")

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	}
}

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

	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(yin.SimpleLogger)
	r.Use(middleware.Recoverer)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("."))
	// })

	//Users
	// r.Route("/", func(r chi.Router) {
	// 	r.Get("/dashboard", isAuthorized(controllers.Dashboard(users)))
	// 	r.Post("/login", isAuthorized(controllers.Login(users)))
	// 	r.Post("/callback", isAuthorized(controllers.Callback(users)))
	// 	r.Delete("/logout", isAuthorized(controllers.Logout(users)))

	// 	r.Route("/profile", func(r chi.Router) {
	// 		r.Get("/", isAuthorized(controllers.GetUser(users)))
	// 		r.Put("/", isAuthorized(controllers.UpdateUser(users)))
	// 	})
	// })

	//Announcements
	r.Route("/announcements", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetAnnouncements(announcements)))
		r.Post("/", isAuthorized(controllers.AddAnnouncement(announcements)))

		r.Route("/{announcementID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetAnnouncement(announcements)))
			r.Put("/", isAuthorized(controllers.UpdateAnnouncement(announcements)))
			r.Delete("/", isAuthorized(controllers.DeleteAnnouncement(announcements)))
		})
	})

	//Events
	r.Route("/events", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetEvents(events)))
		r.Post("/", isAuthorized(controllers.AddEvent(events)))

		r.Route("/{eventID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetEvent(events)))
			r.Put("/", isAuthorized(controllers.UpdateEvent(events)))
			r.Delete("/", isAuthorized(controllers.DeleteEvent(events)))
		})
	})

	//Holidays
	r.Route("/holidays", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetHolidays(holidays)))
		r.Post("/", isAuthorized(controllers.AddHoliday(holidays)))

		r.Route("/{holidayID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetHoliday(holidays)))
			r.Put("/", isAuthorized(controllers.UpdateHoliday(holidays)))
			r.Delete("/", isAuthorized(controllers.DeleteHoliday(holidays)))
		})
	})

	//Members
	r.Route("/members", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetMembers(members)))
		r.Post("/", isAuthorized(controllers.AddMember(members)))

		r.Route("/{memberID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetMember(members)))
			r.Put("/", isAuthorized(controllers.UpdateMember(members)))
			r.Delete("/", isAuthorized(controllers.DeleteMember(members)))
		})
	})

	//Organizations
	r.Route("/organizations", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetOrganizations(organizations)))
		r.Post("/", isAuthorized(controllers.AddOrganization(organizations)))

		r.Route("/{organizationID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetOrganization(organizations)))
			r.Put("/", isAuthorized(controllers.UpdateOrganization(organizations)))
			r.Delete("/", isAuthorized(controllers.DeleteOrganization(organizations)))
		})
	})

	//Scriptures
	r.Route("/scriptures", func(r chi.Router) {
		r.Get("/", isAuthorized(controllers.GetScriptures(scriptures)))
		r.Post("/", isAuthorized(controllers.AddScripture(scriptures)))

		r.Route("/{scriptureID}", func(r chi.Router) {
			r.Get("/", isAuthorized(controllers.GetScripture(scriptures)))
			r.Put("/", isAuthorized(controllers.UpdateScripture(scriptures)))
			r.Delete("/", isAuthorized(controllers.DeleteScripture(scriptures)))

			//Nested Chapters
			r.Route("/chapters", func(r chi.Router) {
				r.Get("/", isAuthorized(controllers.GetChapters(chapters)))
				r.Post("/", isAuthorized(controllers.AddChapter(chapters)))

				r.Route("/{chapterID}", func(r chi.Router) {
					r.Get("/", isAuthorized(controllers.GetChapter(chapters)))
					r.Put("/", isAuthorized(controllers.UpdateChapter(chapters)))
					r.Delete("/", isAuthorized(controllers.DeleteChapter(chapters)))
				})
			})
		})
	})

	// Mount the admin sub-router, which btw is the same as:
	// r.Mount("/admin", adminRouter())

	http.ListenAndServe(":3000", r)
}

// A completely separate router for administrator routes
// func adminRouter() chi.Router {
// 	r := chi.NewRouter()
// 	r.Use(AdminOnly)
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("admin: dashboard"))
// 	})
// 	r.Get("/account", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("admin: show account."))
// 	})
// 	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(fmt.Sprintf("admin: view user id %v", chi.URLParam(r, "userId"))))
// 	})
// 	return r
// }

// func AdminOnly(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		isAdmin, ok := r.Context().Value("acl.admin").(bool)
// 		if !ok || !isAdmin {
// 			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
