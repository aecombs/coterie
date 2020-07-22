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
		CREATE TABLE IF NOT EXISTS event (
			ID	INT NOT NULL UNIQUE AUTO_INCREMENT,
			name	TEXT,
			occasion	TEXT,
			date	TEXT,
			description	TEXT,
			created_at	TEXT,
			updated_at	TEXT,
			organization_id	INT,
			PRIMARY KEY(ID),
			FOREIGN KEY(organization_id) REFERENCES organization(ID)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &EventTable{
		DB: db,
	}
}

//EventsLister grabs all the events for an org
func (eventTable *EventTable) EventsLister(orgID string) ([]Event, error) {
	events := []Event{}
	rows, err := eventTable.DB.Query(`
		SELECT * FROM event WHERE event.organization_id = ?
	`, orgID)
	if err != nil {
		log.Printf("Unable to retrieve events: %s", err.Error())
		return nil, err
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
	return events, nil
}

//Model.where(id: "")
func (eventTable *EventTable) EventGetter(eventID string) (Event, error) {
	var event Event

	stmt, err := eventTable.DB.Prepare(`
		SELECT * FROM event WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Event{}, err
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
			log.Printf("Unable to retrieve event: %s", err.Error())
			return Event{}, err
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
	return event, nil
}

//Model.create
func (eventTable *EventTable) EventAdder(event Event) (Event, error) {
	stmt, err := eventTable.DB.Prepare(`
		INSERT INTO event (name,occasion,date,description,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?,?)
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Event{}, err
	}

	_, err = stmt.Exec(event.Name, event.Occasion, event.Date, event.Description, event.CreatedAt, event.UpdatedAt, event.OrganizationID)

	if err != nil {
		log.Printf("Unable to add event: %s", err.Error())
		return Event{}, err
	}
	defer stmt.Close()

	return event, nil
}

//Model.update
func (eventTable *EventTable) EventUpdater(event Event) (Event, error) {
	stmt, err := eventTable.DB.Prepare(`
	UPDATE event SET name = ?, occasion = ?, date = ?, description = ?, updated_at = ? WHERE event.id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Event{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Occasion, event.Date, event.Description, event.UpdatedAt, event.ID)

	if err != nil {
		log.Printf("Unable to update event: %s", err.Error())
		return Event{}, err
	}
	return event, nil
}

//Model.delete
func (eventTable *EventTable) EventDeleter(eventID string) error {
	stmt, err := eventTable.DB.Prepare(`
		DELETE FROM event WHERE event.id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventID)

	if err != nil {
		log.Printf("Unable to delete event: %s", err.Error())
		return err
	}

	return nil
}
