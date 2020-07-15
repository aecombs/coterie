package controllers

import (
	"coterie/models"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/qkgo/yin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Data struct {
	ID       string
	Name     string
	Email    string
	Picture  string
	Verified bool
	Primary  bool
}

var (
	googleOauthConfig *oauth2.Config

	//randomized string of nums
	oauthStateString = "oauth" + strconv.Itoa(rand.Intn(999999-111111)+111111)
)

func init() {
	//seed value
	rand.Seed(time.Now().UnixNano())

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     goDotEnvVariable("GOOGLE_CLIENT_ID"),
		ClientSecret: goDotEnvVariable("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

//Google Login
func GoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := googleOauthConfig.AuthCodeURL(oauthStateString)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

//GoogleCallback is the callback func that receives Google's data after an Oauth request is approved. It will then add a new user if they did not already exist in the system and redirects the user to their dashboard.
func GoogleCallback(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		response, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
		if err != nil {
			log.Printf("Unable to retrieve user info from Google: %s", err.Error())
			res.SendStatus(400)
			// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		//logic to check if they exit in database.
		user, err := AddUser(userTable, response)
		if err != nil {
			log.Printf("Unable to register user: %s", err.Error())
			res.SendStatus(404)
			return
		}
		log.Printf("%s", user)

		//request cookie:
		cookie, err := r.Cookie("session")
		//it it doesn't exist, we receive an err. Set the cookie!
		if err != nil {
			//TODO: refactor to use uuid
			sessionID := strconv.Itoa(user.ID)
			cookie = &http.Cookie{
				Name:     "session",
				Value:    sessionID,
				HttpOnly: true,
				Path:     "/",
				Expires:  time.Now().Add(time.Hour * 24 * 14),
			}
			http.SetCookie(w, cookie)
		}

		url := "http://localhost:3001/dashboard"
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	}
}

//GetUserInfo is a helper func for GoogleCallback. It parses the info and returns it as a Data struct
func getUserInfo(state string, code string) (Data, error) {
	if state != oauthStateString {
		log.Println("Invalid oauth state")
		return Data{}, errors.New("Invalid Oauth State")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Code exchange failed: %s", err.Error())
		return Data{}, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("Failed getting user info: %s", err.Error())
		return Data{}, err
	}

	defer response.Body.Close()

	data := Data{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Printf("Unable to decode into json: %s", err.Error())
		log.Fatal(err)
	}

	return data, nil
}

//AddUser Checks if a user exists in system, creating one if it doesn't already exist
//Returns User instance
func AddUser(userTable *models.UserTable, content Data) (models.User, error) {
	userBefore := models.User{
		GoogleID:  content.ID,
		Name:      content.Name,
		Email:     content.Email,
		Bio:       "",
		Avatar:    content.Picture,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	//check google_id against db
	existingUser, err := userTable.UserGetterByGoogleID(userBefore.GoogleID)

	//if err is nil, that means we either retrieved the user or they do not exist in the database
	if err != nil {
		log.Printf("Unable to retrieve existing user from database: %s", err.Error())
		return models.User{}, err
	}
	if existingUser.Name == userBefore.Name {
		return existingUser, nil
	}
	//if we get here, that means this is the first time the user has logged in and we need to save their info
	newUser, err := userTable.RegisterUser(userBefore)
	if err != nil {
		log.Printf("Unable to add user to database: %s", err.Error())
		return models.User{}, err
	}
	return newUser, nil
}

//LogoutUser will change the session cookie so that it no longer contains the userID.
func LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		//it it doesn't exist, we receive an err. No need to delete anything.
		if err == nil {
			//reset the cookie to have a "deleted" value and to expire immediately
			cookie = &http.Cookie{
				Value:   "deleted",
				Expires: time.Now(),
			}
			http.SetCookie(w, cookie)
		}

		url := "/"
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return
	}
}

//GetUser returns a single instance of User
func GetUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		user, err := GrabLoggedInUser(userTable, r)
		if err != nil {
			log.Printf("Unable to grab user: %s", err.Error())
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(user)
	}
}

//UpdateUser will update the name, email, and bio
func UpdateUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		user, err := GrabLoggedInUser(userTable, r)
		if err != nil {
			log.Printf("Unable to grab user: %s", err.Error())
		}
		body := map[string]string{}
		req.BindBody(&body)

		updatedUser := models.User{
			ID:        user.ID,
			Name:      body["name"],
			Email:     body["email"],
			Bio:       body["bio"],
			UpdatedAt: time.Now().String(),
		}

		result, err := userTable.UserUpdater(updatedUser)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendJSON(result)
	}
}
