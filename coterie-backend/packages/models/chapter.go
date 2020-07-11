package models

import (
	"database/sql"
)

type Chapter struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Text        string `json:"text,omitempty"`
	Position    int    `json:"position,omitempty"`
	ScriptureID int    `json:"scripture_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type ChapterTable struct {
	DB *sql.DB
}

func NewChapterTable(db *sql.DB) *ChapterTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "chapter" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"text"  TEXT,
			"position"  INTEGER,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"scripture_id"	INTEGER,
			FOREIGN KEY("scripture_id") REFERENCES "scripture"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()
	return &ChapterTable{
		DB: db,
	}
}

// //Model.All
// func (announcementTable *AnnouncementTable) AnnouncementsLister() ([]Announcement, error) {
// 	announcements := []Announcement{}
// 	rows, err := announcementTable.DB.Query(`
// 		SELECT * FROM announcement
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var id int
// 	var text string
// 	var date string
// 	var createdAt string
// 	var updatedAt string
// 	var organizationID int
// 	for rows.Next() {
// 		rows.Scan(&id, &text, &date, &createdAt, &updatedAt, &organizationID)
// 		announcement := Announcement{
// 			ID:             id,
// 			Text:           text,
// 			Date:           date,
// 			CreatedAt:      createdAt,
// 			UpdatedAt:      updatedAt,
// 			OrganizationID: organizationID,
// 		}
// 		announcements = append(announcements, announcement)
// 	}
// 	return announcements, err
// }

// //Model.where(id: "")
// func (announcementTable *AnnouncementTable) AnnouncementGetter(announcementID string) (Announcement, error) {
// 	var announcement Announcement

// 	stmt, err := announcementTable.DB.Prepare(`
// 		SELECT * FROM announcement WHERE id = ?
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	if stmt != nil {
// 		var id int
// 		var text string
// 		var date string
// 		var createdAt string
// 		var updatedAt string
// 		var organizationID int

// 		err = stmt.QueryRow(announcementID).Scan(&id, &text, &date, &createdAt, &updatedAt, &organizationID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		announcement.ID = id
// 		announcement.Text = text
// 		announcement.Date = date
// 		announcement.CreatedAt = createdAt
// 		announcement.UpdatedAt = updatedAt
// 		announcement.OrganizationID = organizationID
// 	}
// 	return announcement, err
// }

// //Model.create
// func (announcementTable *AnnouncementTable) AnnouncementAdder(announcement Announcement) (Announcement, error) {
// 	stmt, err := announcementTable.DB.Prepare(`
// 		INSERT INTO announcement (date,text,created_at,updated_at,organization_id) values (?,?,?,?,?)
// 	`)

// 	stmt.Exec(announcement.Date, announcement.Text, announcement.CreatedAt, announcement.UpdatedAt, announcement.OrganizationID)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	return announcement, err
// }

// //Model.update
// func (announcementTable *AnnouncementTable) AnnouncementUpdater(announcement Announcement) (Announcement, error) {
// 	stmt, err := announcementTable.DB.Prepare(`
// 	UPDATE announcement SET date = ?, text = ?, updated_at = ? WHERE announcement.id = ?
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(announcement.Date, announcement.Text, announcement.UpdatedAt, announcement.ID)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return announcement, err
// }

// //Model.delete
// func (announcementTable *AnnouncementTable) AnnouncementDeleter(announcementID string) error {
// 	stmt, err := announcementTable.DB.Prepare(`
// 		DELETE FROM announcement WHERE announcement.id = ?
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(announcementID)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return err
// }
