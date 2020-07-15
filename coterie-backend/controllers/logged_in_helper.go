package controllers

import (
	"coterie/models"
	"log"
	"net/http"
)

//LoggedInUser checks the cookie key Session to see if it's valid and exists. If it does, it returns the user associated with it.
func LoggedInUser(userTable *models.UserTable, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("No session cookie stored in cookies: %s", err.Error())
		return models.User{}, err
	}
	user, err := userTable.UserGetterByID(cookie.Value)
	if err != nil {
		log.Printf("Unable to retrieve user: %s", err.Error())
		return models.User{}, err
	}
	return user, nil
}
