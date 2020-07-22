package models

import (
	"database/sql"
	"log"
	"strconv"
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
		CREATE TABLE IF NOT EXISTS organization (
			ID	INT NOT NULL UNIQUE AUTO_INCREMENT,
			name	TEXT,
			mission_statement	TEXT,
			total_funds	INT,
			created_at	TEXT,
			updated_at	TEXT,
			user_id	INT,
			PRIMARY KEY(ID),
			FOREIGN KEY(user_id) REFERENCES user(ID)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &OrganizationTable{
		DB: db,
	}
}

//OrganizationsLister grabs all the Organizations for a user
func (organizationTable *OrganizationTable) OrganizationsLister(usID string) ([]Organization, error) {
	organizations := []Organization{}
	rows, err := organizationTable.DB.Query(`
		SELECT * FROM organization WHERE user_id = ?
	`, usID)
	if err != nil {
		log.Printf("Unable to retrieve organizations: %s", err.Error())
		return nil, err
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

		updatedFunds, err := organizationTable.updateFunds(id)
		if err != nil {
			log.Printf("Bad Query: %s", err.Error())
			return nil, err
		}

		organization := Organization{
			ID:               id,
			Name:             name,
			MissionStatement: missionStatement,
			TotalFunds:       updatedFunds,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
			UserID:           userID,
		}
		organizations = append(organizations, organization)
	}
	return organizations, nil
}

//Model.where(id: "")
func (organizationTable *OrganizationTable) OrganizationGetter(organizationID string) (Organization, error) {
	var organization Organization

	stmt, err := organizationTable.DB.Prepare(`
		SELECT * FROM organization WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Organization{}, err
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var missionStatement string
		var totalFunds int
		var createdAt string
		var updatedAt string
		var userID int

		err = stmt.QueryRow(organizationID).Scan(&id, &name, &missionStatement, &totalFunds, &createdAt, &updatedAt, &userID)
		if err != nil {
			log.Printf("Unable to retrieve organization: %s", err.Error())
			return Organization{}, err
		}

		updatedFunds, err := organizationTable.updateFunds(id)
		if err != nil {
			log.Printf("Bad Query: %s", err.Error())
			return Organization{}, err
		}
		log.Printf("UpdatedFunds:: %s", strconv.Itoa(updatedFunds))

		organization.ID = id
		organization.Name = name
		organization.MissionStatement = missionStatement
		organization.TotalFunds = updatedFunds
		organization.CreatedAt = createdAt
		organization.UpdatedAt = updatedAt
		organization.UserID = userID
	}
	return organization, nil
}

//Model.create
func (organizationTable *OrganizationTable) OrganizationAdder(organization Organization) (Organization, error) {
	stmt, err := organizationTable.DB.Prepare(`
		INSERT INTO organization (name,mission_statement,total_funds,created_at,updated_at,user_id) VALUES (?,?,?,?,?,?)
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Organization{}, err
	}

	stmt.Exec(organization.Name, organization.MissionStatement, organization.TotalFunds, organization.CreatedAt, organization.UpdatedAt, organization.UserID)

	if err != nil {
		log.Printf("Unable to create org: %s", err.Error())
		return Organization{}, err
	}
	defer stmt.Close()

	return organization, nil
}

//Model.update
func (organizationTable *OrganizationTable) OrganizationUpdater(organization Organization) (Organization, error) {
	stmt, err := organizationTable.DB.Prepare(`
	UPDATE organization SET name = ?, mission_statement = ?, total_funds = ?, updated_at = ? WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Organization{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(organization.Name, organization.MissionStatement, organization.TotalFunds, organization.UpdatedAt, organization.ID)

	if err != nil {
		log.Printf("Unable to update org: %s", err.Error())
		return Organization{}, err
	}
	return organization, nil
}

//Model.delete
func (organizationTable *OrganizationTable) OrganizationDeleter(organizationID string) error {
	stmt, err := organizationTable.DB.Prepare(`
		DELETE FROM organization WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(organizationID)

	if err != nil {
		log.Printf("Unable to delete org: %s", err.Error())
		return err
	}

	return nil
}

func (organizationTable *OrganizationTable) updateFunds(orgID int) (int, error) {
	var calculatedFunds int

	rows, err := organizationTable.DB.Query(`
		SELECT funds_raised FROM member WHERE member.organization_id = ?
	`, strconv.Itoa(orgID))
	if err != nil {
		log.Printf("Unable to retrieve information: %s", err.Error())
		return 0, err
	}
	defer rows.Close()

	var fundsRaised int

	for rows.Next() {
		rows.Scan(&fundsRaised)
		calculatedFunds += fundsRaised
	}

	return calculatedFunds, nil

}
