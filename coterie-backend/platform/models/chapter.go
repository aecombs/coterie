package models

import "database/sql"

type Chapter struct {
	DB *sql.DB
}

func NewChapterTable(db *sql.DB) *Chapter {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "chapter" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"text"  TEXT,
			"position"  INTEGER,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Chapter{
		DB: db,
	}
}
