package models

import "database/sql"

type Member struct {
	DB *sql.DB
}

func NewMemberTable(db *sql.DB) *Member {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "member" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"birthdate"  DATE,
			"class"  TEXT,
			"email"  TEXT,
			"funds_raised"  INTEGER,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &Member{
		DB: db,
	}
}
