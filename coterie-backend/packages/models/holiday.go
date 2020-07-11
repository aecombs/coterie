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
		CREATE TABLE IF NOT EXISTS "holiday" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"date"  TEXT,
			"description"  TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &HolidayTable{
		DB: db,
	}
}

//Model.All
func (holidayTable *HolidayTable) HolidaysLister() ([]Holiday, error) {
	holidays := []Holiday{}
	rows, err := holidayTable.DB.Query(`
		SELECT * FROM holiday
	`)
	if err != nil {
		log.Fatal(err)
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
	return holidays, err
}

//Model.where(id: "")
func (holidayTable *HolidayTable) HolidayGetter(holidayID string) (Holiday, error) {
	var holiday Holiday

	stmt, err := holidayTable.DB.Prepare(`
		SELECT * FROM holiday WHERE id = ?
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

		err = stmt.QueryRow(holidayID).Scan(&id, &name, &date, &description, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		holiday.ID = id
		holiday.Name = name
		holiday.Date = date
		holiday.Description = description
		holiday.CreatedAt = createdAt
		holiday.UpdatedAt = updatedAt
		holiday.OrganizationID = organizationID
	}
	return holiday, err
}

//Model.create
func (holidayTable *HolidayTable) HolidayAdder(holiday Holiday) (Holiday, error) {
	stmt, err := holidayTable.DB.Prepare(`
		INSERT INTO holiday (name,date,description,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(holiday.Name, holiday.Date, holiday.Description, holiday.CreatedAt, holiday.UpdatedAt, holiday.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return holiday, err
}

//Model.update
func (holidayTable *HolidayTable) HolidayUpdater(holiday Holiday) (Holiday, error) {
	stmt, err := holidayTable.DB.Prepare(`
	UPDATE holiday SET name = ?, date = ?, description = ?, updated_at = ? WHERE holiday.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(holiday.Name, holiday.Date, holiday.Description, holiday.UpdatedAt, holiday.ID)

	if err != nil {
		log.Fatal(err)
	}
	return holiday, err
}

//Model.delete
func (holidayTable *HolidayTable) HolidayDeleter(holidayID string) error {
	stmt, err := holidayTable.DB.Prepare(`
		DELETE FROM holiday WHERE holiday.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(holidayID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
