package models

import (
	"database/sql"
	"log"
)

type Announcement struct {
	ID             int    `json:"id,omitempty"`
	Text           string `json:"text,omitempty"`
	Date           string `json:"date,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type AnnouncementTable struct {
	DB *sql.DB
}

func NewAnnouncementTable(db *sql.DB) *AnnouncementTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "announcement" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"text"	TEXT,
			"date"	TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &AnnouncementTable{
		DB: db,
	}
}

//Model.All
func (announcementTable *AnnouncementTable) AnnouncementsLister() []Announcement {
	announcements := []Announcement{}
	rows, _ := announcementTable.DB.Query(`
		SELECT * FROM announcement
	`)
	var id int
	var text string
	var date string
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &text, &date, &createdAt, &updatedAt, &organizationID)
		announcement := Announcement{
			ID:             id,
			Text:           text,
			Date:           date,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		announcements = append(announcements, announcement)
	}
	return announcements
}

//Model.where(id: "")
func (announcementTable *AnnouncementTable) AnnouncementGetter(id int) *Announcement {
	// var announcement *Announcement
	rows, _ := announcementTable.DB.Query("SELECT * FROM announcement WHERE id = ?", id)

	announcement := Announcement{
		ID:             id,
		Text:           text,
		Date:           date,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		OrganizationID: organizationID,
	}

	return announcement
}

//Model.create
func (announcementTable *AnnouncementTable) AnnouncementAdder(announcement Announcement) int {
	stmt, err := announcementTable.DB.Prepare(`
		INSERT INTO announcement (date,text,created_at,updated_at) values (?,?,?,?)
	`)

	stmt.Exec(announcement.Date, announcement.Text, announcement.CreatedAt, announcement.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		// res.SendStatus(400)
		return 400
	} else {
		return 204
		// res.SendJSON(announcement)
	}
}
