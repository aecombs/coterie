package models

import "database/sql"

type Event struct {
	DB *sql.DB
}

func NewEventTable(db *sql.DB) *Event {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "event" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"date"	DATE,
			"type"	TEXT,
			"description"	TEXT,
			PRIMARY KEY("ID")
		);
	`)

	stmt.Exec()
	return &Event{
		DB: db,
	}
}
