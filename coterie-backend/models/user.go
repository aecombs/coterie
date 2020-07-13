package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
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
func (userTable *UserTable) UserGetter(option string, userID string) (User, error) {
	var user User

	stmt, err := userTable.DB.Prepare(`
		SELECT * FROM user WHERE ? = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var name string
		var email string
		var bio string
		var avatar string
		var createdAt string
		var updatedAt string

		err = stmt.QueryRow(option, userID).Scan(&id, &name, &email, &bio, &avatar, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.Bio = bio
		user.Avatar = avatar
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	return user, err
}

//Model.create...only used when user is new
func (userTable *UserTable) UserAdder(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
		INSERT INTO user (name,email,bio,avatar,created_at,updated_at) VALUES (?,?,?,?,?,?)
	`)

	stmt.Exec(user.Name, user.Email, user.Bio, user.Avatar, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return user, err
}

//Model.update
func (userTable *UserTable) UserUpdater(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
	UPDATE user SET name = ?, email = ?, bio = ?, updated_at = ? WHERE user.id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Bio, user.UpdatedAt, user.ID)

	if err != nil {
		log.Fatal(err)
	}
	return user, err
}
