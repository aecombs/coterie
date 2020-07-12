package controllers

import (
	"coterie/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

// // AdminOnly middleware restricts access to just administrators.
// func AdminOnly(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		isAdmin, ok := r.Context().Value("acl.admin").(bool)
// 		if !ok || !isAdmin {
// 			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

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
