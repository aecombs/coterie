package controllers

import (
	"coterie/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//Show
func GetUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		userID := chi.URLParam(r, "userID")

		user, err := userTable.UserGetter(userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(user)
	}
}

//Create
func AddUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)

		user := models.User{
			Name:      body["name"],
			Email:     body["email"],
			Avatar:    body["avatar"],
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		result, err := userTable.Login(user)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//Update
func UpdateUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		userID := chi.URLParam(r, "userID")
		body := map[string]string{}
		req.BindBody(&body)

		IntUserID, _ := strconv.Atoi(userID)
		funds, _ := strconv.Atoi(body["funds_raised"])
		user := models.User{
			ID:        IntUserID,
			Name:      body["name"],
			Email:     body["email"],
			UpdatedAt: time.Now().String(),
		}

		result, err := userTable.UserUpdater(user)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}
