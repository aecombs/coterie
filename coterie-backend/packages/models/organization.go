package models

import (
	"database/sql"
	"log"
)

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

//Model.All
func (organizationTable *OrganizationTable) OrganizationsLister() ([]Organization, error) {
	organizations := []Organization{}
	rows, err := organizationTable.DB.Query(`
		SELECT * FROM organization
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string
	var missionStatement string
	var totalFunds int
	var createdAt string
	var updatedAt string
	var userID int

	for rows.Next() {
		rows.Scan(&id, &name, &missionStatement, &totalFunds, &createdAt, &updatedAt, &userID)
		organization := Organization{
			ID:               id,
			Name:             name,
			MissionStatement: missionStatement,
			TotalFunds:       totalFunds,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			UserID:           userID,
		}
		organizations = append(organizations, organization)
	}
	return organizations, err
}

//Model.where(id: "")
func (organizationTable *OrganizationTable) OrganizationGetter(organizationID string) (Organization, error) {
	var organization Organization

	stmt, err := organizationTable.DB.Prepare(`
		SELECT * FROM organization WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var date string
		var description string
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(organizationID).Scan(&id, &name, &date, &description, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		organization.ID = id
		organization.Name = name
		organization.Date = date
		organization.Description = description
		organization.CreatedAt = createdAt
		organization.UpdatedAt = updatedAt
		organization.OrganizationID = organizationID
	}
	return organization, err
}

//Model.create
func (organizationTable *OrganizationTable) OrganizationAdder(organization Organization) (Organization, error) {
	stmt, err := organizationTable.DB.Prepare(`
		INSERT INTO organization (name,date,description,created_at,updated_at,user_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(organization.Name, organization.Date, organization.Description, organization.CreatedAt, organization.UpdatedAt, organization.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return organization, err
}

//Model.update
func (organizationTable *OrganizationTable) OrganizationUpdater(organization Organization) (Organization, error) {
	stmt, err := organizationTable.DB.Prepare(`
	UPDATE organization SET name = ?, date = ?, description = ?, updated_at = ? WHERE organization.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(organization.Name, organization.Date, organization.Description, organization.UpdatedAt, organization.ID)

	if err != nil {
		log.Fatal(err)
	}
	return organization, err
}

//Model.delete
func (organizationTable *OrganizationTable) OrganizationDeleter(organizationID string) error {
	stmt, err := organizationTable.DB.Prepare(`
		DELETE FROM organization WHERE organization.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(organizationID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
