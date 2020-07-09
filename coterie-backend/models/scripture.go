package models

import "database/sql"

type Scripture struct {
	DB *sql.DB
}

func NewScriptureTable(db *sql.DB) *Scripture {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "scripture" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Scripture{
		DB: db,
	}
}
