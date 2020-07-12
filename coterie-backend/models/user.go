package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
}

type UserTable struct {
	DB *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "user" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT,
			"email"  TEXT,
			"bio"  TEXT,
			"avatar"  TEXT,
			"created_at"	TEXT,
			"updated_at"	TEXT,
			"organization_id"	INTEGER,
			FOREIGN KEY("organization_id") REFERENCES "organization"("ID"),
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	defer stmt.Close()

	return &UserTable{
		DB: db,
	}
}

//Model.where(id: "")
func (userTable *UserTable) UserGetter(userID string) (User, error) {
	var user User

	stmt, err := userTable.DB.Prepare(`
		SELECT * FROM user WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var email string
		var avatar string
		var createdAt string
		var updatedAt string
		var organizationID int

		err = stmt.QueryRow(userID).Scan(&id, &name, &email, &avatar, &createdAt, &updatedAt, &organizationID)
		if err != nil {
			log.Fatal(err)
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.Avatar = avatar
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
		user.OrganizationID = organizationID
	}
	return user, err
}

//Model.login
func (userTable *UserTable) Login(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
		INSERT INTO user (name,email,avatar,created_at,updated_at,organization_id) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(user.Name, user.Email, user.Avatar, user.CreatedAt, user.UpdatedAt, user.OrganizationID)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return user, err
}

//Model.update
func (userTable *UserTable) UserUpdater(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
	UPDATE user SET name = ?, email = ?, avatar = ?, updated_at = ? WHERE user.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Avatar, user.UpdatedAt, user.ID)

	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
