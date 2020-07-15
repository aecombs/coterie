package controllers

import (
	"coterie/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//goDotEnvVariable grabs a value from the .env file
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//GrabLoggedInUser checks the cookie key Session to see if it's valid and exists. If it does, it returns the user associated with it.
func GrabLoggedInUser(userTable *models.UserTable, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil && cookie.Value == "deleted" {
		log.Printf("User is not logged in: %s", err.Error())
		return models.User{}, err
	}
	user, err := userTable.UserGetterByID(cookie.Value)
	if err != nil {
		log.Printf("Unable to retrieve user: %s", err.Error())
		return models.User{}, err
	}
	return user, nil
}
