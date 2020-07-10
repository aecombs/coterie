package models

import "database/sql"

type Holiday struct {
	DB *sql.DB
}

func NewHolidayTable(db *sql.DB) *Holiday {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "holiday" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"date"  DATE,
			"description"  TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Holiday{
		DB: db,
	}
}
