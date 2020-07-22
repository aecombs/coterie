package main

import (
	"coterie/controllers"
	"coterie/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/qkgo/yin"
)

func main() {

	// If the optional DB_TCP_HOST environment variable is set, it contains
	// the IP address and port number of a TCP connection pool to be created,
	// such as "127.0.0.1:3306". If DB_TCP_HOST is not set, a Unix socket
	// connection pool will be created instead.

	//Access the cloud SQL database
	// var (
	// 	err                    = godotenv.Load(".env")
	// 	dbUser                 = os.Getenv("DB_USER")
	// 	dbPwd                  = os.Getenv("DB_PASS")
	// 	instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	// 	dbName                 = os.Getenv("DB_NAME")
	// )

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "cloudsql"
	// }

	// var dbURI string

	// dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	//access the database
	// db, err := sql.Open("mysql", dbURI)
	// if err != nil {
	// 	log.Printf("Unable to access database: %s", err.Error())
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	db, err := initSocketConnectionPool()
	if err != nil {
		log.Fatalf("initSocketConnectionPool: unable to connect: %v", err)
	}

	users := models.NewUserTable(db)
	organizations := models.NewOrganizationTable(db)
	announcements := models.NewAnnouncementTable(db)
	scriptures := models.NewScriptureTable(db)
	chapters := models.NewChapterTable(db)
	events := models.NewEventTable(db)
	holidays := models.NewHolidayTable(db)
	members := models.NewMemberTable(db)

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

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func getEnvOrDie(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_socket]
	var (
		dbUser                 = getEnvOrDie("DB_USER")
		dbPwd                  = getEnvOrDie("DB_PASS")
		instanceConnectionName = getEnvOrDie("INSTANCE_CONNECTION_NAME")
		dbName                 = getEnvOrDie("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_mysql_databasesql_create_socket]
}

// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_mysql_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_mysql_databasesql_limit]

	// [START cloud_sql_mysql_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_mysql_databasesql_lifetime]
}
