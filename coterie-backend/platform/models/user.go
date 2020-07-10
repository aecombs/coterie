package models

import "database/sql"

type User struct {
	DB *sql.DB
}

func NewUserTable(db *sql.DB) *User {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "user" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"email"  TEXT,
			"avatar"  TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &User{
		DB: db,
	}
}
