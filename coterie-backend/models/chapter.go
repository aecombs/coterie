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

	stmt.Exec()

	defer stmt.Close()

	return &ChapterTable{
		DB: db,
	}
}

//Model.All
func (chapterTable *ChapterTable) ChaptersLister(scripID string) ([]Chapter, error) {
	chapters := []Chapter{}
	rows, err := chapterTable.DB.Query(`
		SELECT * FROM chapter WHERE chapter.scripture_id = ?
	`, scripID)
	if err != nil {
		log.Printf("Unable to retrieve chapters: %s", err.Error())
		return nil, err
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
	return chapters, nil
}

//Model.where(id: "")
func (chapterTable *ChapterTable) ChapterGetter(chapterID string) (Chapter, error) {
	var chapter Chapter

	stmt, err := chapterTable.DB.Prepare(`
		SELECT * FROM chapter WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Chapter{}, err
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
			log.Printf("Unable to retrieve chapter: %s", err.Error())
			return Chapter{}, err
		}

		chapter.ID = id
		chapter.Name = name
		chapter.Text = text
		chapter.Position = position
		chapter.CreatedAt = createdAt
		chapter.UpdatedAt = updatedAt
		chapter.ScriptureID = scriptureID
	}
	return chapter, nil
}

//Model.create
func (chapterTable *ChapterTable) ChapterAdder(chapter Chapter) (Chapter, error) {
	stmt, err := chapterTable.DB.Prepare(`
		INSERT INTO chapter (name,text,position,created_at,updated_at,scripture_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(chapter.Name, chapter.Text, chapter.Position, chapter.CreatedAt, chapter.UpdatedAt, chapter.ScriptureID)

	if err != nil {
		log.Printf("Unable to create chapter: %s", err.Error())
		return Chapter{}, err
	}
	defer stmt.Close()

	return chapter, nil
}

//Model.update
func (chapterTable *ChapterTable) ChapterUpdater(chapter Chapter) (Chapter, error) {
	stmt, err := chapterTable.DB.Prepare(`
	UPDATE chapter SET name = ?, text = ?, position = ?, updated_at = ? WHERE chapter.id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Chapter{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chapter.Name, chapter.Text, chapter.Position, chapter.UpdatedAt, chapter.ID)

	if err != nil {
		log.Printf("Unable to update chapter: %s", err.Error())
		return Chapter{}, err
	}
	return chapter, nil
}

//Model.delete
func (chapterTable *ChapterTable) ChapterDeleter(chapterID string) error {
	stmt, err := chapterTable.DB.Prepare(`
		DELETE FROM chapter WHERE chapter.id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chapterID)

	if err != nil {
		log.Printf("Unable to delete chapter: %s", err.Error())
		return err
	}

	return nil
}
