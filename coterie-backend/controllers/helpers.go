package controllers

import (
	"coterie/models"
	"fmt"
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
	cookie, err := r.Cookie("__session")

	if err == nil && (fmt.Sprintf("%T", cookie.Value) != "int") || err != nil {
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

//EnableCors alters the Header to allow Cross Origin Resource Sharing
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}
