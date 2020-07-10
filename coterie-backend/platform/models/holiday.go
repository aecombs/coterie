package models

import "database/sql"

type Holiday struct {
	ID          int
	Name        string
	Date        string
	Description string
	created_at  string
	updated_at  string
}

type HolidayTable struct {
	DB *sql.DB
}

func NewHolidayTable(db *sql.DB) *HolidayTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "holiday" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"date"  DATE,
			"description"  TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &HolidayTable{
		DB: db,
	}
}
