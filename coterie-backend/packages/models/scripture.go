package models

import "database/sql"

type Scripture struct {
	ID             int
	Name           string
	OrganizationID int
	CreatedAt      string
	UpdatedAt      string
}

type ScriptureTable struct {
	DB *sql.DB
}

func NewScriptureTable(db *sql.DB) *ScriptureTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "scripture" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &ScriptureTable{
		DB: db,
	}
}
