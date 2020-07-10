package models

import (
	"database/sql"
)

type Announcement struct {
	ID         int
	Text       string
	Date       string
	created_at string
	updated_at string
}

type AnnouncementTable struct {
	DB *sql.DB
}

func NewAnnouncementTable(db *sql.DB) *AnnouncementTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "announcement" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"text"	TEXT,
			"date"	DATE,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &AnnouncementTable{
		DB: db,
	}
}

// func ListAllAnnouncements() []Announcement {
// 	announcements := NewAnnouncementTable(Database)
// 	return announcements
// }
