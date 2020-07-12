package controllers

import (
	"coterie/packages/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//Show
func GetUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		// userID := chi.URLParam(r, "userID")
		

		user, err := userTable.UserGetter(userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(user)
	}
}
