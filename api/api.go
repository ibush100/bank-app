package api

import (
	"bank-app/helpers"
	"bank-app/interfaces"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	user := interfaces.User{UserID: userID, Username: username, Email: email, Password: password}
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{})
	db.Create(&user)
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
	w.WriteHeader(http.StatusCreated)
	helpers.WriteToJson(w, formattedBody)
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
