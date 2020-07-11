package models

import (
	"database/sql"
	"log"
)

type Chapter struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Text        string `json:"text,omitempty"`
	Position    int    `json:"position,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	ScriptureID int    `json:"scripture_id,omitempty"`
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

	defer stmt.Close()

	stmt.Exec()

	return &ChapterTable{
		DB: db,
	}
}

//Model.All
func (chapterTable *ChapterTable) ChaptersLister() ([]Chapter, error) {
	chapters := []Chapter{}
	rows, err := chapterTable.DB.Query(`
		SELECT * FROM chapter
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string
	var text string
	var position int
	var createdAt string
	var updatedAt string
	var scriptureID int
	for rows.Next() {
		rows.Scan(&id, &name, &text, &position, &createdAt, &updatedAt, &scriptureID)
		chapter := Chapter{
			ID:          id,
			Name:        name,
			Text:        text,
			Position:    position,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			ScriptureID: scriptureID,
		}
		chapters = append(chapters, chapter)
	}
	return chapters, err
}

//Model.where(id: "")
func (chapterTable *ChapterTable) ChapterGetter(chapterID string) (Chapter, error) {
	var chapter Chapter

	stmt, err := chapterTable.DB.Prepare(`
		SELECT * FROM chapter WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var text string
		var position int
		var createdAt string
		var updatedAt string
		var scriptureID int

		err = stmt.QueryRow(chapterID).Scan(&id, &name, &text, &position, &createdAt, &updatedAt, &scriptureID)
		if err != nil {
			log.Fatal(err)
		}

		chapter.ID = id
		chapter.Name = name
		chapter.Text = text
		chapter.Position = position
		chapter.CreatedAt = createdAt
		chapter.UpdatedAt = updatedAt
		chapter.ScriptureID = scriptureID
	}
	return chapter, err
}

//Model.create
func (chapterTable *ChapterTable) ChapterAdder(chapter Chapter) (Chapter, error) {
	stmt, err := chapterTable.DB.Prepare(`
		INSERT INTO chapter (name,text,position,created_at,updated_at,scripture_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(chapter.Name, chapter.Text, chapter.Position, chapter.CreatedAt, chapter.UpdatedAt, chapter.ScriptureID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return chapter, err
}

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
