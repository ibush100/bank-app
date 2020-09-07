package api

import (
	"bank-app/helpers"
	"bank-app/interfaces"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func createUser(username string, email string, password string) (interfaces.User, bool) {
	userID := uuid.Must(uuid.NewRandom())
	passwordHash := helpers.HashAndSalt([]byte(password))
	user := interfaces.User{UserID: userID, Username: username, Email: email, Password: passwordHash}
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{})
	db.Create(&user)
	// need to clean up returning true
	return user, true
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var fomattedUser interfaces.Register
	err := json.Unmarshal(body, &fomattedUser)
	helpers.HandleErr(err)
	registerUser, result := createUser(fomattedUser.Username, fomattedUser.Email, fomattedUser.Password)
	if result {
		w.WriteHeader(http.StatusCreated)
		helpers.WriteToJson(w, registerUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Call successful dude")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody interfaces.User
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	if isUserPresent(formattedBody.Email, formattedBody.Password) {
		token := PrepareToken(formattedBody.ID)
		w.WriteHeader(http.StatusCreated)
		// create token
		helpers.WriteToJson(w, token)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func isUserPresent(email string, password string) bool {
	userResult := FindUser(email)
	if userResult <= 0 {
		return false
	}
	return true
}

func FindUser(email string) uint {
	db := helpers.ConnectDB()
	var user interfaces.User
	//db.Table("users").Select("user_id").Where("email = ? ", email).First(&user.ID)
	db.Where("email = ?", email).First(&user)

	return user.ID
}

func PrepareToken(ID uint) string {
	tokenContent := jwt.MapClaims{
		"user_id": ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/", index)
	router.HandleFunc("/login", registerUser).Methods("POST")
	router.HandleFunc("/register", registerUser).Methods("POST")
	fmt.Println("App is working on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
