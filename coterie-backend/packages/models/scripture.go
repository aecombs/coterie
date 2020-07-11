package models

import (
	"database/sql"
	"log"
)

type Scripture struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type ScriptureTable struct {
	DB *sql.DB
}

func NewScriptureTable(db *sql.DB) *ScriptureTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "scripture" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &ScriptureTable{
		DB: db,
	}
}

//Model.All
func (scriptureTable *ScriptureTable) ScripturesLister() ([]Scripture, error) {
	scriptures := []Scripture{}
	rows, err := scriptureTable.DB.Query(`
		SELECT * FROM scripture
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &name, &createdAt, &updatedAt, &organizationID)
		scripture := Scripture{
			ID:             id,
			Name:           name,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		scriptures = append(scriptures, scripture)
	}
	return scriptures, err
}

//Model.where(id: "")
func (scriptureTable *ScriptureTable) ScriptureGetter(scriptureID string) (Scripture, error) {
	var scripture Scripture

	stmt, err := scriptureTable.DB.Prepare(`
		SELECT * FROM scripture WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(scriptureID).Scan(&id, &name, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		scripture.ID = id
		scripture.Name = name
		scripture.CreatedAt = createdAt
		scripture.UpdatedAt = updatedAt
		scripture.OrganizationID = organizationID
	}
	return scripture, err
}

//Model.create
func (scriptureTable *ScriptureTable) ScriptureAdder(scripture Scripture) (Scripture, error) {
	stmt, err := scriptureTable.DB.Prepare(`
		INSERT INTO scripture (name,created_at,updated_at,organization_id) VALUES (?,?,?,?)
	`)

	stmt.Exec(scripture.Name, scripture.CreatedAt, scripture.UpdatedAt, scripture.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return scripture, err
}

//Model.update
func (scriptureTable *ScriptureTable) ScriptureUpdater(scripture Scripture) (Scripture, error) {
	stmt, err := scriptureTable.DB.Prepare(`
	UPDATE scripture SET name = ?, updated_at = ? WHERE scripture.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(scripture.Name, scripture.UpdatedAt, scripture.ID)

	if err != nil {
		log.Fatal(err)
	}
	return scripture, err
}

//Model.delete
func (scriptureTable *ScriptureTable) ScriptureDeleter(scriptureID string) error {
	stmt, err := scriptureTable.DB.Prepare(`
		DELETE FROM scripture WHERE scripture.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(scriptureID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
