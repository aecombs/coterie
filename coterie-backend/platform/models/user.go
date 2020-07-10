package models

import "database/sql"

type User struct {
	ID         int
	Name       string
	Email      string
	Avatar     string
	created_at string
	updated_at string
}

type UserTable struct {
	DB *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
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
	return &UserTable{
		DB: db,
	}
}
