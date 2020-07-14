package controllers

import (
	"coterie/models"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
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

func LoadEnv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		//TODO: update this to use session
		data := goDotEnvVariable("GOOGLE_CLIENT_SECRET")

		res.SendJSON(data)
	}
}

// var mySigningKey = goDotEnvVariable("MY_JWT_TOKEN")

// func GenerateJWT() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["user"] = "Your Mom"
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Errorf("Something went wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }
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

//Google Callback
func GoogleCallback(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		response, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
		if err != nil {
			fmt.Println(err.Error())
			res.SendStatus(400)
			// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		//logic to check if they exit in database.
		user, err := AddUser(userTable, response)
		if err != nil {
			fmt.Println(err.Error())
			res.SendStatus(404)
			// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		//send content
		// if they do, return the info that react needs.
		// if they don't, create a new user and return the info react needs.
		//Note: React might only need the userID to store in it's session

		fmt.Fprintf(w, "Responded with: %s\n", response)

		res.SendJSON(user)
	}
}

func getUserInfo(state string, code string) (Data, error) {
	if state != oauthStateString {
		fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	data := Data{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

//Create New User
func AddUser(userTable *models.UserTable, content Data) (models.User, error) {
	//logic to part content
	userBefore := models.User{
		GoogleID:  content.ID,
		Name:      content.Name,
		Email:     content.Email,
		Bio:       "",
		Avatar:    content.Picture,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	_, err := userTable.RegisterUser(userBefore)
	if err != nil {
		fmt.Errorf("Unable to add user to database")
	}

	userAfter, err := userTable.UserGetter("email", userBefore.Email)
	if err != nil {
		fmt.Errorf("Something went wrong")
	}
	// userID := strconv.Itoa(userAfter.ID)

	return userBefore, nil
}

//Logout
func LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: reset session...Or is that React's job?
		url := "/"
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return
	}
}

//Show
func GetUser(userTable *models.UserTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		//TODO: update this to use session
		userID := chi.URLParam(r, "userID")

		user, err := userTable.UserGetter("id", userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(user)
	}
}

// func AddUser(userTable *models.UserTable) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		res, req := yin.Event(w, r)
// 		body := map[string]string{}
// 		req.BindBody(&body)

// 		user := models.User{
// 			Name:      body["name"],
// 			Email:     body["email"],
// 			Avatar:    body["avatar"],
// 			CreatedAt: time.Now().String(),
// 			UpdatedAt: time.Now().String(),
// 		}

// 		result, err := userTable.UserAdder(user)
// 		if err != nil {
// 			http.Error(w, http.StatusText(404), 404)
// 			return
// 		}

// 		res.SendJSON(result)
// 	}
// }

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
