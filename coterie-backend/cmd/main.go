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
	users := models.NewUserTable(db)

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
	r.Route("/", func(r chi.Router) {
		// 	r.Get("/dashboard", controllers.Dashboard(users))
		r.Get("/auth/google", controllers.GoogleLogin())
		r.Get("/auth/google/callback", controllers.GoogleCallback(users))
		r.Delete("/logout", controllers.LogoutUser())

		r.Route("/profile", func(r chi.Router) {
			r.Get("/", controllers.GetUser(users))
			r.Put("/", controllers.UpdateUser(users))
		})
	})

	//Announcements
	r.Route("/announcements", func(r chi.Router) {
		r.Get("/", controllers.GetAnnouncements(announcements))
		r.Post("/", controllers.AddAnnouncement(announcements))

		r.Route("/{announcementID}", func(r chi.Router) {
			r.Get("/", controllers.GetAnnouncement(announcements))
			r.Put("/", controllers.UpdateAnnouncement(announcements))
			r.Delete("/", controllers.DeleteAnnouncement(announcements))
		})
	})

	//Events
	r.Route("/events", func(r chi.Router) {
		r.Get("/", controllers.GetEvents(events))
		r.Post("/", controllers.AddEvent(events))

		r.Route("/{eventID}", func(r chi.Router) {
			r.Get("/", controllers.GetEvent(events))
			r.Put("/", controllers.UpdateEvent(events))
			r.Delete("/", controllers.DeleteEvent(events))
		})
	})

	//Holidays
	r.Route("/holidays", func(r chi.Router) {
		r.Get("/", controllers.GetHolidays(holidays))
		r.Post("/", controllers.AddHoliday(holidays))

		r.Route("/{holidayID}", func(r chi.Router) {
			r.Get("/", controllers.GetHoliday(holidays))
			r.Put("/", controllers.UpdateHoliday(holidays))
			r.Delete("/", controllers.DeleteHoliday(holidays))
		})
	})

	//Members
	r.Route("/members", func(r chi.Router) {
		r.Get("/", controllers.GetMembers(members))
		r.Post("/", controllers.AddMember(members))

		r.Route("/{memberID}", func(r chi.Router) {
			r.Get("/", controllers.GetMember(members))
			r.Put("/", controllers.UpdateMember(members))
			r.Delete("/", controllers.DeleteMember(members))
		})
	})

	//Organizations
	r.Route("/organizations", func(r chi.Router) {
		r.Get("/", controllers.GetOrganizations(organizations))
		r.Post("/", controllers.AddOrganization(organizations))

		r.Route("/{organizationID}", func(r chi.Router) {
			r.Get("/", controllers.GetOrganization(organizations))
			r.Put("/", controllers.UpdateOrganization(organizations))
			r.Delete("/", controllers.DeleteOrganization(organizations))
		})
	})

	//Scriptures
	r.Route("/scriptures", func(r chi.Router) {
		r.Get("/", controllers.GetScriptures(scriptures))
		r.Post("/", controllers.AddScripture(scriptures))

		r.Route("/{scriptureID}", func(r chi.Router) {
			r.Get("/", controllers.GetScripture(scriptures))
			r.Put("/", controllers.UpdateScripture(scriptures))
			r.Delete("/", controllers.DeleteScripture(scriptures))

			//Nested Chapters
			r.Route("/chapters", func(r chi.Router) {
				r.Get("/", controllers.GetChapters(chapters))
				r.Post("/", controllers.AddChapter(chapters))

				r.Route("/{chapterID}", func(r chi.Router) {
					r.Get("/", controllers.GetChapter(chapters))
					r.Put("/", controllers.UpdateChapter(chapters))
					r.Delete("/", controllers.DeleteChapter(chapters))
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
