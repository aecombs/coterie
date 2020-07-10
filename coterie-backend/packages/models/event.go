package models

import "database/sql"

type Event struct {
	ID          int
	Name        string
	Date        string
	Type        string
	Description string
	created_at  string
	updated_at  string
}

type EventTable struct {
	DB *sql.DB
}

func NewEventTable(db *sql.DB) *EventTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "event" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"date"	DATE,
			"type"	TEXT,
			"description"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID")
		);
	`)

	stmt.Exec()
	return &EventTable{
		DB: db,
	}
}
