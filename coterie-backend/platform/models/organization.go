package models

import "database/sql"

type Organization struct {
	DB *sql.DB
}

func NewOrganizationTable(db *sql.DB) *Organization {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "organization" (
			"ID"	INTEGER NOT NULL,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"mission_statement"	TEXT,
			"total_funds"	INTEGER,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Organization{
		DB: db,
	}
}
