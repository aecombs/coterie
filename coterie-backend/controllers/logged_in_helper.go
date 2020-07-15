package controllers

import (
	"coterie/models"
	"net/http"
)

func LoggedInUser(userTable *models.UserTable, r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return models.User{}, err
	}
	user, err := userTable.UserGetterByID(cookie.Value)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
