package models

import (
	"database/sql"
	"log"
)

type Holiday struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Date           string `json:"date,omitempty"`
	Description    string `json:"description,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type HolidayTable struct {
	DB *sql.DB
}

func NewHolidayTable(db *sql.DB) *HolidayTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS holiday (
			ID	INT NOT NULL UNIQUE AUTO_INCREMENT,
			name	TEXT,
			date  TEXT,
			description  TEXT,
			created_at	TEXT,
			updated_at	TEXT,
			organization_id	INT,
			PRIMARY KEY(ID),
			FOREIGN KEY(organization_id) REFERENCES organization(ID) ON DELETE CASCADE
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &HolidayTable{
		DB: db,
	}
}

//HolidaysLister grabs all the holidays for an org
func (holidayTable *HolidayTable) HolidaysLister(orgID string) ([]Holiday, error) {
	holidays := []Holiday{}
	rows, err := holidayTable.DB.Query(`
		SELECT * FROM holiday WHERE organization_id = ?
	`, orgID)
	if err != nil {
		log.Printf("Unable to retrieve holidays: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var date string
	var description string
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &name, &date, &description, &createdAt, &updatedAt, &organizationID)
		holiday := Holiday{
			ID:             id,
			Name:           name,
			Date:           date,
			Description:    description,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		holidays = append(holidays, holiday)
	}
	return holidays, nil
}

//Model.where(id: "")
func (holidayTable *HolidayTable) HolidayGetter(holidayID string) (Holiday, error) {
	var holiday Holiday

	stmt, err := holidayTable.DB.Prepare(`
		SELECT * FROM holiday WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Holiday{}, err
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

		err = stmt.QueryRow(holidayID).Scan(&id, &name, &date, &description, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Printf("Unable to retrieve holiday: %s", err.Error())
			return Holiday{}, err
		}

		holiday.ID = id
		holiday.Name = name
		holiday.Date = date
		holiday.Description = description
		holiday.CreatedAt = createdAt
		holiday.UpdatedAt = updatedAt
		holiday.OrganizationID = organizationID
	}
	return holiday, nil
}

//Model.create
func (holidayTable *HolidayTable) HolidayAdder(holiday Holiday) (Holiday, error) {
	stmt, err := holidayTable.DB.Prepare(`
		INSERT INTO holiday (name,date,description,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?)
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Holiday{}, err
	}

	_, err = stmt.Exec(holiday.Name, holiday.Date, holiday.Description, holiday.CreatedAt, holiday.UpdatedAt, holiday.OrganizationID)

	if err != nil {
		log.Printf("Unable to create holiday: %s", err.Error())
		return Holiday{}, err
	}
	defer stmt.Close()

	return holiday, nil
}

//Model.update
func (holidayTable *HolidayTable) HolidayUpdater(holiday Holiday) (Holiday, error) {
	stmt, err := holidayTable.DB.Prepare(`
	UPDATE holiday SET name = ?, date = ?, description = ?, updated_at = ? WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Holiday{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(holiday.Name, holiday.Date, holiday.Description, holiday.UpdatedAt, holiday.ID)

	if err != nil {
		log.Printf("Unable to update holiday: %s", err.Error())
		return Holiday{}, err
	}
	return holiday, nil
}

//Model.delete
func (holidayTable *HolidayTable) HolidayDeleter(holidayID string) error {
	stmt, err := holidayTable.DB.Prepare(`
		DELETE FROM holiday WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(holidayID)

	if err != nil {
		log.Printf("Unable to delete holiday: %s", err.Error())
		return err
	}

	return nil
}
