package models

import "database/sql"

type Chapter struct {
	ID         int
	Name       string
	Text       string
	Position   int
	ScriptureIds []*Scripture
	created_at string
	updated_at string
}

type ChapterTable struct {
	DB *sql.DB
}

func NewChapterTable(db *sql.DB) *ChapterTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "chapter" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"text"  TEXT,
			"position"  INTEGER,
			"scripture_id"	INTEGER,
			FOREIGN KEY("scripture_id") REFERENCES "scripture"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &ChapterTable{
		DB: db,
	}
}
