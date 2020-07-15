package models

import (
	"database/sql"
	"log"
	"strconv"
)

//User struct
type User struct {
	ID        int    `json:"id,omitempty"`
	GoogleID  string `json:"google_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

//UserTable struct will be all the users in the DB
type UserTable struct {
	DB *sql.DB
}

//NewUserTable creates a new table when there isn't one that exists already
func NewUserTable(db *sql.DB) *UserTable {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "user" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"google_id" TEXT,
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

//UserGetterByID will retrieve a single instance of user by their user_id
func (userTable *UserTable) UserGetterByID(userID string) (User, error) {
	var user User

	stmt, err := userTable.DB.Prepare(`
			SELECT * FROM user WHERE user.id = ?
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return User{}, err
	}

	defer stmt.Close()

	if stmt != nil {
		var id int
		var googleID string
		var name string
		var email string
		var bio string
		var avatar string
		var createdAt string
		var updatedAt string

		err = stmt.QueryRow(userID).Scan(&id, &googleID, &name, &email, &bio, &avatar, &createdAt, &updatedAt)

		if err == sql.ErrNoRows {
			log.Printf("User does not exist: %s", err.Error())
			return User{}, nil
		} else if err != nil {
			log.Printf("Unable to retrieve user: %s", err.Error())
			return User{}, err
		}

		user.ID = id
		user.GoogleID = googleID
		user.Name = name
		user.Email = email
		user.Bio = bio
		user.Avatar = avatar
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	return user, nil
}

//UserGetterByGoogleID will retrieve a single instance of user by their google_id
func (userTable *UserTable) UserGetterByGoogleID(googleID string) (User, error) {
	var user User

	stmt, err := userTable.DB.Prepare(`
			SELECT * FROM user WHERE user.google_id = ?
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return User{}, err
	}
	defer stmt.Close()

	if stmt != nil {
		var id int
		var gID string
		var name string
		var email string
		var bio string
		var avatar string
		var createdAt string
		var updatedAt string

		err = stmt.QueryRow(googleID).Scan(&id, &gID, &name, &email, &bio, &avatar, &createdAt, &updatedAt)

		if err == sql.ErrNoRows {
			log.Printf("User does not exist: %s", err.Error())
			return User{}, nil
		} else if err != nil {
			log.Printf("Unable to retrieve user: %s", err.Error())
			return User{}, err
		}

		user.ID = id
		user.GoogleID = gID
		user.Name = name
		user.Email = email
		user.Bio = bio
		user.Avatar = avatar
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	return user, nil
}

//RegisterUser will create a new user when they log in for the first time
func (userTable *UserTable) RegisterUser(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
		INSERT INTO user (google_id,name,email,bio,avatar,created_at,updated_at) VALUES (?,?,?,?,?,?,?)
	`)

	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return User{}, err
	}

	_, err = stmt.Exec(user.GoogleID, user.Name, user.Email, user.Bio, user.Avatar, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		log.Printf("Unable to create user: %s", err.Error())
		return User{}, err
	}

	//retrieve their newly created ID from DB
	var id string

	err = userTable.DB.QueryRow("SELECT id FROM user WHERE google_id = ?", user.GoogleID).Scan(&id)

	if err != nil {
		log.Printf("Unable to retrieve user ID from database: %s", err.Error())
		return User{}, err
	}

	defer stmt.Close()

	user.ID, err = strconv.Atoi(id)

	return user, nil
}

//UserUpdater will update a single instance of user in the database
func (userTable *UserTable) UserUpdater(user User) (User, error) {
	stmt, err := userTable.DB.Prepare(`
	UPDATE user SET name = ?, email = ?, bio = ?, updated_at = ? WHERE user.id = ?
	`)
	if err != nil {
		log.Printf("Bad Query: %s", err.Error())
		return User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Bio, user.UpdatedAt, user.ID)

	if err != nil {
		log.Printf("Unable to update user: %s", err.Error())
		return User{}, err
	}
	return user, nil
}
