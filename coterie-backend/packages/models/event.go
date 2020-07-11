package models

import (
	"database/sql"
	"log"
)

type Event struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Occasion       string `json:"occasion,omitempty"`
	Date           string `json:"date,omitempty"`
	Description    string `json:"description,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type EventTable struct {
	DB *sql.DB
}

func NewEventTable(db *sql.DB) *EventTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "event" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"occasion"	TEXT,
			"date"	TEXT,
			"description"	TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID")
		);
	`)

	stmt.Exec()
	return &EventTable{
		DB: db,
	}
}

//Model.All
func (eventTable *EventTable) EventsLister() ([]Event, error) {
	events := []Event{}
	rows, err := eventTable.DB.Query(`
		SELECT * FROM event
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string
	var occasion string
	var date string
	var description string
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &name, &occasion, &date, &description, &createdAt, &updatedAt, &organizationID)
		event := Event{
			ID:             id,
			Name:           name,
			Occasion:       occasion,
			Date:           date,
			Description:    description,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		events = append(events, event)
	}
	return events, err
}

//Model.where(id: "")
func (eventTable *EventTable) EventGetter(eventID string) (Event, error) {
	var event Event

	stmt, err := eventTable.DB.Prepare(`
		SELECT * FROM event WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var occasion string
		var date string
		var description string
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(eventID).Scan(&id, &name, &occasion, &date, &description, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		event.ID = id
		event.Name = name
		event.Occasion = occasion
		event.Date = date
		event.Description = description
		event.CreatedAt = createdAt
		event.UpdatedAt = updatedAt
		event.OrganizationID = organizationID
	}
	return event, err
}

//Model.create
func (eventTable *EventTable) EventAdder(event Event) (Event, error) {
	stmt, err := eventTable.DB.Prepare(`
		INSERT INTO event (name,occasion,date,description,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(event.Name, event.Occasion, event.Date, event.Description, event.CreatedAt, event.UpdatedAt, event.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return event, err
}

//Model.update
func (eventTable *EventTable) EventUpdater(event Event) (Event, error) {
	stmt, err := eventTable.DB.Prepare(`
	UPDATE event SET name = ?, occasion = ?, date = ?, description = ?, updated_at = ? WHERE event.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Occasion, event.Date, event.Description, event.UpdatedAt, event.ID)

	if err != nil {
		log.Fatal(err)
	}
	return event, err
}

//Model.delete
func (eventTable *EventTable) EventDeleter(eventID string) error {
	stmt, err := eventTable.DB.Prepare(`
		DELETE FROM event WHERE event.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
