package models

import "database/sql"

type Member struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Birthdate      string `json:"birthdate,omitempty"`
	Class          string `json:"class,omitempty"`
	Email          string `json:"email,omitempty"`
	FundsRaised    int    `json:"funds_raised,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type MemberTable struct {
	DB *sql.DB
}

func NewMemberTable(db *sql.DB) *MemberTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "member" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"birthdate"  TEXT,
			"class"  TEXT,
			"email"  TEXT,
			"funds_raised"  INTEGER,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &MemberTable{
		DB: db,
	}
}
