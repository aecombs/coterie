package models

import "database/sql"

type Scripture struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type ScriptureTable struct {
	DB *sql.DB
}

func NewScriptureTable(db *sql.DB) *ScriptureTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "scripture" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &ScriptureTable{
		DB: db,
	}
}
