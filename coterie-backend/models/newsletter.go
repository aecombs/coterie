package models

import (
	"database/sql"
	"log"
)

type Newsletter struct {
	ID             int    `json:"id,omitempty"`
	Header         string `json:"header,omitempty"`
	Description    string `json:"description,omitempty"`
	Date           string `json:"date,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type NewsletterTable struct {
	DB *sql.DB
}

func NewNewsletterTable(db *sql.DB) *NewsletterTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS newsletter (
			ID INT NOT NULL UNIQUE AUTO_INCREMENT,
			header	TEXT,
			description	 TEXT,
			date  TEXT,
			created_at  TEXT,
			updated_at  TEXT,
			organization_id INT,
			PRIMARY KEY(ID),
			FOREIGN KEY (organization_id) REFERENCES organization(ID) ON DELETE CASCADE
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &NewsletterTable{
		DB: db,
	}
}

//NewslettersLister grabs all the newsletters for an org.
func (newsletterTable *NewsletterTable) NewslettersLister(orgID string) ([]Newsletter, error) {
	newsletters := []Newsletter{}
	rows, err := newsletterTable.DB.Query(`
		SELECT * FROM newsletter WHERE organization_id = ?
	`, orgID)
	if err != nil {
		log.Printf("Unable to retrieve newsletters: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var id int
	var header string
	var description string
	var date string
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &header, &description, &date, &createdAt, &updatedAt, &organizationID)
		newsletter := Newsletter{
			ID:             id,
			Header:         header,
			Description:    description,
			Date:           date,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		newsletters = append(newsletters, newsletter)
	}
	return newsletters, nil
}

//Model.where(id: "")
func (newsletterTable *NewsletterTable) NewsletterGetter(newsletterID string) (Newsletter, error) {
	var newsletter Newsletter

	stmt, err := newsletterTable.DB.Prepare(`
		SELECT * FROM newsletter WHERE id = ?
	`)
	if err != nil {
		log.Printf("Invalid sql query: %s", err.Error())
		return Newsletter{}, err
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var header string
		var description string
		var date string
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(newsletterID).Scan(&id, &header, &description, &date, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Printf("Unable to retrieve newsletter: %s", err.Error())
			return Newsletter{}, err
		}

		newsletter.ID = id
		newsletter.Header = header
		newsletter.Description = description
		newsletter.Date = date
		newsletter.CreatedAt = createdAt
		newsletter.UpdatedAt = updatedAt
		newsletter.OrganizationID = organizationID
	}
	return newsletter, nil
}

//Model.create
func (newsletterTable *NewsletterTable) NewsletterAdder(newsletter Newsletter) (Newsletter, error) {
	stmt, err := newsletterTable.DB.Prepare(`
		INSERT INTO newsletter (header,description,date,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?)
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Newsletter{}, err
	}

	_, err = stmt.Exec(newsletter.Header, newsletter.Description, newsletter.Date, newsletter.CreatedAt, newsletter.UpdatedAt, newsletter.OrganizationID)

	if err != nil {
		log.Printf("Unable to create newsletter: %s", err.Error())
		return Newsletter{}, err
	}
	defer stmt.Close()

	return newsletter, nil
}

//Model.update
func (newsletterTable *NewsletterTable) NewsletterUpdater(newsletter Newsletter) (Newsletter, error) {
	stmt, err := newsletterTable.DB.Prepare(`
	UPDATE newsletter SET header = ?, description = ?, date = ?, updated_at = ? WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return Newsletter{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newsletter.Header, newsletter.Description, newsletter.Date, newsletter.UpdatedAt, newsletter.ID)

	if err != nil {
		log.Printf("Unable to update newsletter: %s", err.Error())
		return Newsletter{}, err
	}
	return newsletter, nil
}

//Model.delete
func (newsletterTable *NewsletterTable) NewsletterDeleter(newsletterID string) error {
	stmt, err := newsletterTable.DB.Prepare(`
		DELETE FROM newsletter WHERE id = ?
	`)
	if err != nil {
		log.Printf("Bad query: %s", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newsletterID)

	if err != nil {
		log.Printf("Unable to delete newsletter: %s", err.Error())
		return err
	}

	return nil
}
