package models

import "database/sql"

type Organization struct {
	ID          int
	Name        string
	Date        string
	Type        string
	Description string
	UserID      int
	CreatedAt   string
	UpdatedAt   string
}

type OrganizationTable struct {
	DB *sql.DB
}

func NewOrganizationTable(db *sql.DB) *OrganizationTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "organization" (
			"ID"	INTEGER NOT NULL,
			"created_at"	DATETIME,
			"updated_at"	DATETIME,
			"name"	TEXT,
			"mission_statement"	TEXT,
			"total_funds"	INTEGER,
			"user_id"	INTEGER,
			FOREIGN KEY("user_id") REFERENCES "user"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &OrganizationTable{
		DB: db,
	}
}
