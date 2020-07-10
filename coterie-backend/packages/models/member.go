package models

import "database/sql"

type Member struct {
	ID          int
	Name        string
	Birthdate   string
	Class       string
	Email       string
	FundsRaised int
	created_at  string
	updated_at  string
}

type MemberTable struct {
	DB *sql.DB
}

func NewMemberTable(db *sql.DB) *MemberTable {
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
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &MemberTable{
		DB: db,
	}
}
