package models

import (
	"database/sql"
	"log"
)

type Member struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Birthdate      string `json:"birthdate,omitempty"`
	Class          string `json:"class,omitempty"`
	Email          string `json:"email,omitempty"`
	FundsRaised    int    `json:"funds_raised,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type MemberTable struct {
	DB *sql.DB
}

func NewMemberTable(db *sql.DB) *MemberTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "member" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"birthdate"  TEXT,
			"class"  TEXT,
			"email"  TEXT,
			"funds_raised"  INTEGER,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &MemberTable{
		DB: db,
	}
}

//Model.All
func (memberTable *MemberTable) MembersLister() ([]Member, error) {
	members := []Member{}
	rows, err := memberTable.DB.Query(`
		SELECT * FROM member
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string
	var birthdate string
	var class string
	var email string
	var fundsRaised int
	var createdAt string
	var updatedAt string
	var organizationID int
	for rows.Next() {
		rows.Scan(&id, &name, &birthdate, &class, &email, &fundsRaised, &createdAt, &updatedAt, &organizationID)
		member := Member{
			ID:             id,
			Name:           name,
			Birthdate:      birthdate,
			Class:          class,
			Email:          email,
			FundsRaised:    fundsRaised,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			OrganizationID: organizationID,
		}
		members = append(members, member)
	}
	return members, err
}

//Model.where(id: "")
func (memberTable *MemberTable) MemberGetter(memberID string) (Member, error) {
	var member Member

	stmt, err := memberTable.DB.Prepare(`
		SELECT * FROM member WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var birthdate string
		var class string
		var email string
		var fundsRaised int
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(memberID).Scan(&id, &name, &birthdate, &class, &email, &fundsRaised, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		member.ID = id
		member.Name = name
		member.Birthdate = birthdate
		member.Class = class
		member.Email = email
		member.FundsRaised = fundsRaised
		member.CreatedAt = createdAt
		member.UpdatedAt = updatedAt
		member.OrganizationID = organizationID
	}
	return member, err
}

//Model.create
func (memberTable *MemberTable) MemberAdder(member Member) (Member, error) {
	stmt, err := memberTable.DB.Prepare(`
		INSERT INTO member (name, birthdate, class, email, funds_raised, created_at, updated_at, organization_id) VALUES (?,?,?,?,?,?,?,?)
	`)

	stmt.Exec(member.Name, member.Birthdate, member.Class, member.Email, member.FundsRaised, member.CreatedAt, member.UpdatedAt, member.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return member, err
}

//Model.update
func (memberTable *MemberTable) MemberUpdater(member Member) (Member, error) {
	stmt, err := memberTable.DB.Prepare(`
	UPDATE member SET class = ?, email = ?, funds_raised = ?, updated_at = ? WHERE member.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(member.Class, member.Email, member.FundsRaised, member.UpdatedAt, member.ID)

	if err != nil {
		log.Fatal(err)
	}
	return member, err
}

//Model.delete
func (memberTable *MemberTable) MemberDeleter(memberID string) error {
	stmt, err := memberTable.DB.Prepare(`
		DELETE FROM member WHERE member.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(memberID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
