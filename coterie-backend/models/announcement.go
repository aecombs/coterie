package models

import "database/sql"

type Announcement struct {
	DB *sql.DB
}

func NewAnnouncementTable(db *sql.DB) *Announcement {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "announcement" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"text"	TEXT,
			"date"	DATE,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Announcement{
		DB: db,
	}
}
