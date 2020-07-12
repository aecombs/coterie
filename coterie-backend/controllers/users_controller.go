package controllers

import (
	"coterie/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/qkgo/yin"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var mySigningKey = goDotEnvVariable("MY_JWT_TOKEN")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Your Mom"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// var (
// 	googleOauthConfig *oauth2.Config
// )

// func init() {
// 	googleOauthConfig = &oauth2.Config{
// 		RedirectURL:  "http://localhost:8080/callback",
// 		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
// 		Endpoint:     google.Endpoint,
// 	}
// }

//Login
//Callback
//Logout

//Show
func GetUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		//TODO: update this to use session
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

		result, err := userTable.UserAdder(user)
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
		//TODO: Update this to use session
		userID := chi.URLParam(r, "userID")
		body := map[string]string{}
		req.BindBody(&body)

		IntUserID, _ := strconv.Atoi(userID)
		user := models.User{
			ID:        IntUserID,
			Name:      body["name"],
			Email:     body["email"],
			Bio:       body["bio"],
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
