package models

import "database/sql"

type Organization struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	MissionStatement string `json:"mission_statement,omitempty"`
	TotalFunds       int    `json:"total_funds,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UserID           int    `json:"user_id,omitempty"`
}

type OrganizationTable struct {
	DB *sql.DB
}

func NewOrganizationTable(db *sql.DB) *OrganizationTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "organization" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"mission_statement"	TEXT,
			"total_funds"	INTEGER,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"user_id"	INTEGER,
			FOREIGN KEY("user_id") REFERENCES "user"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &OrganizationTable{
		DB: db,
	}
}
